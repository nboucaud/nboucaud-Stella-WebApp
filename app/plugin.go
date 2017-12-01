// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	l4g "github.com/alecthomas/log4go"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/utils"

	builtinplugin "github.com/mattermost/mattermost-server/app/plugin"
	"github.com/mattermost/mattermost-server/app/plugin/jira"
	"github.com/mattermost/mattermost-server/app/plugin/ldapextras"

	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/plugin/pluginenv"
)

var prepackagedPlugins map[string]func(string) ([]byte, error) = map[string]func(string) ([]byte, error){
	"jira": jira.Asset,
}

func (a *App) initBuiltInPlugins() {
	plugins := map[string]builtinplugin.Plugin{
		"ldapextras": &ldapextras.Plugin{},
	}
	for id, p := range plugins {
		l4g.Info("Initializing plugin: " + id)
		api := &BuiltInPluginAPI{
			id:     id,
			router: a.Srv.Router.PathPrefix("/plugins/" + id).Subrouter(),
			app:    a,
		}
		p.Initialize(api)
	}
	utils.AddConfigListener(func(before, after *model.Config) {
		for _, p := range plugins {
			p.OnConfigurationChange()
		}
	})
	for _, p := range plugins {
		p.OnConfigurationChange()
	}
}

// ActivatePlugins will activate any plugins enabled in the config
// and deactivate all other plugins.
func (a *App) ActivatePlugins() {
	if a.PluginEnv == nil {
		l4g.Error("plugin env not initialized")
		return
	}

	plugins, err := a.PluginEnv.Plugins()
	if err != nil {
		l4g.Error("failed to activate plugins: " + err.Error())
		return
	}

	for _, plugin := range plugins {
		id := plugin.Manifest.Id

		pluginState := &model.PluginState{Enable: false}
		if state, ok := a.Config().PluginSettings.PluginStates[id]; ok {
			pluginState = state
		}

		active := a.PluginEnv.IsPluginActive(id)

		if pluginState.Enable && !active {
			if err := a.PluginEnv.ActivatePlugin(id); err != nil {
				l4g.Error(err.Error())
				continue
			}

			if plugin.Manifest.HasClient() {
				message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_PLUGIN_ACTIVATED, "", "", "", nil)
				message.Add("manifest", plugin.Manifest.ClientManifest())
				a.Publish(message)
			}

			l4g.Info("Activated %v plugin", id)
		} else if !pluginState.Enable && active {
			if err := a.PluginEnv.DeactivatePlugin(id); err != nil {
				l4g.Error(err.Error())
				continue
			}

			if plugin.Manifest.HasClient() {
				message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_PLUGIN_DEACTIVATED, "", "", "", nil)
				message.Add("manifest", plugin.Manifest.ClientManifest())
				a.Publish(message)
			}

			l4g.Info("Deactivated %v plugin", id)
		}
	}
}

// InstallPlugin unpacks and installs a plugin but does not activate it.
func (a *App) InstallPlugin(pluginFile io.Reader) (*model.Manifest, *model.AppError) {
	return a.installPlugin(pluginFile, false)
}

func (a *App) installPlugin(pluginFile io.Reader, allowPrepackaged bool) (*model.Manifest, *model.AppError) {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return nil, model.NewAppError("installPlugin", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	tmpDir, err := ioutil.TempDir("", "plugintmp")
	if err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.filesystem.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	defer os.RemoveAll(tmpDir)

	if err := utils.ExtractTarGz(pluginFile, tmpDir); err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.extract.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	tmpPluginDir := tmpDir
	dir, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.filesystem.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	if len(dir) == 1 && dir[0].IsDir() {
		tmpPluginDir = filepath.Join(tmpPluginDir, dir[0].Name())
	}

	manifest, _, err := model.FindManifest(tmpPluginDir)
	if err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.manifest.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	if _, ok := prepackagedPlugins[manifest.Id]; ok && !allowPrepackaged {
		return nil, model.NewAppError("installPlugin", "app.plugin.prepackaged.app_error", nil, "", http.StatusBadRequest)
	}

	bundles, err := a.PluginEnv.Plugins()
	if err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.install.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	for _, bundle := range bundles {
		if bundle.Manifest.Id == manifest.Id {
			return nil, model.NewAppError("installPlugin", "app.plugin.install_id.app_error", nil, "", http.StatusBadRequest)
		}
	}

	err = utils.CopyDir(tmpPluginDir, filepath.Join(a.PluginEnv.SearchPath(), manifest.Id))
	if err != nil {
		return nil, model.NewAppError("installPlugin", "app.plugin.mvdir.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	// Should add manifest validation and error handling here

	return manifest, nil
}

func (a *App) GetPlugins() (*model.PluginsResponse, *model.AppError) {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return nil, model.NewAppError("GetPlugins", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	plugins, err := a.PluginEnv.Plugins()
	if err != nil {
		return nil, model.NewAppError("GetPlugins", "app.plugin.get_plugins.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	resp := &model.PluginsResponse{Active: []*model.PluginInfo{}, Inactive: []*model.PluginInfo{}}
	for _, plugin := range plugins {
		info := &model.PluginInfo{
			Manifest: *plugin.Manifest,
		}
		_, info.Prepackaged = prepackagedPlugins[plugin.Manifest.Id]
		if a.PluginEnv.IsPluginActive(plugin.Manifest.Id) {
			resp.Active = append(resp.Active, info)
		} else {
			resp.Inactive = append(resp.Inactive, info)
		}
	}

	return resp, nil
}

func (a *App) GetActivePluginManifests() ([]*model.Manifest, *model.AppError) {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return nil, model.NewAppError("GetActivePluginManifests", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	plugins := a.PluginEnv.ActivePlugins()

	manifests := make([]*model.Manifest, len(plugins))
	for i, plugin := range plugins {
		manifests[i] = plugin.Manifest
	}

	return manifests, nil
}

func (a *App) RemovePlugin(id string) *model.AppError {
	return a.removePlugin(id, false)
}

func (a *App) removePlugin(id string, allowPrepackaged bool) *model.AppError {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return model.NewAppError("removePlugin", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	if _, ok := prepackagedPlugins[id]; ok && !allowPrepackaged {
		return model.NewAppError("removePlugin", "app.plugin.prepackaged.app_error", nil, "", http.StatusBadRequest)
	}

	plugins, err := a.PluginEnv.Plugins()
	if err != nil {
		return model.NewAppError("removePlugin", "app.plugin.deactivate.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	var manifest *model.Manifest
	for _, p := range plugins {
		if p.Manifest.Id == id {
			manifest = p.Manifest
			break
		}
	}

	if manifest == nil {
		return model.NewAppError("removePlugin", "app.plugin.not_installed.app_error", nil, "", http.StatusBadRequest)
	}

	if a.PluginEnv.IsPluginActive(id) {
		err := a.PluginEnv.DeactivatePlugin(id)
		if err != nil {
			return model.NewAppError("removePlugin", "app.plugin.deactivate.app_error", nil, err.Error(), http.StatusBadRequest)
		}

		if manifest.HasClient() {
			message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_PLUGIN_DEACTIVATED, "", "", "", nil)
			message.Add("manifest", manifest.ClientManifest())
			a.Publish(message)
		}
	}

	err = os.RemoveAll(filepath.Join(a.PluginEnv.SearchPath(), id))
	if err != nil {
		return model.NewAppError("removePlugin", "app.plugin.remove.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// EnablePlugin will set the config for an installed plugin to enabled, triggering activation if inactive.
func (a *App) EnablePlugin(id string) *model.AppError {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return model.NewAppError("EnablePlugin", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	plugins, err := a.PluginEnv.Plugins()
	if err != nil {
		return model.NewAppError("EnablePlugin", "app.plugin.config.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	var manifest *model.Manifest
	for _, p := range plugins {
		if p.Manifest.Id == id {
			manifest = p.Manifest
			break
		}
	}

	if manifest == nil {
		return model.NewAppError("EnablePlugin", "app.plugin.not_installed.app_error", nil, "", http.StatusBadRequest)
	}

	a.UpdateConfig(func(cfg *model.Config) {
		cfg.PluginSettings.PluginStates[id] = &model.PluginState{Enable: true}
	})

	if err := a.SaveConfig(a.Config(), true); err != nil {
		return model.NewAppError("EnablePlugin", "app.plugin.config.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// DisablePlugin will set the config for an installed plugin to disabled, triggering deactivation if active.
func (a *App) DisablePlugin(id string) *model.AppError {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		return model.NewAppError("DisablePlugin", "app.plugin.disabled.app_error", nil, "", http.StatusNotImplemented)
	}

	plugins, err := a.PluginEnv.Plugins()
	if err != nil {
		return model.NewAppError("DisablePlugin", "app.plugin.config.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	var manifest *model.Manifest
	for _, p := range plugins {
		if p.Manifest.Id == id {
			manifest = p.Manifest
			break
		}
	}

	if manifest == nil {
		return model.NewAppError("DisablePlugin", "app.plugin.not_installed.app_error", nil, "", http.StatusBadRequest)
	}

	a.UpdateConfig(func(cfg *model.Config) {
		cfg.PluginSettings.PluginStates[id] = &model.PluginState{Enable: false}
	})

	if err := a.SaveConfig(a.Config(), true); err != nil {
		return model.NewAppError("DisablePlugin", "app.plugin.config.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (a *App) InitPlugins(pluginPath, webappPath string) {
	if !*a.Config().PluginSettings.Enable {
		return
	}

	if a.PluginEnv != nil {
		return
	}

	l4g.Info("Starting up plugins")

	if err := os.Mkdir(pluginPath, 0744); err != nil && !os.IsExist(err) {
		l4g.Error("failed to start up plugins: " + err.Error())
		return
	}

	if err := os.Mkdir(webappPath, 0744); err != nil && !os.IsExist(err) {
		l4g.Error("failed to start up plugins: " + err.Error())
		return
	}

	if env, err := pluginenv.New(
		pluginenv.SearchPath(pluginPath),
		pluginenv.WebappPath(webappPath),
		pluginenv.APIProvider(func(m *model.Manifest) (plugin.API, error) {
			return &PluginAPI{
				id:  m.Id,
				app: a,
				keyValueStore: &PluginKeyValueStore{
					id:  m.Id,
					app: a,
				},
			}, nil
		}),
	); err != nil {
		l4g.Error("failed to start up plugins: " + err.Error())
		return
	} else {
		a.PluginEnv = env
	}

	for id, asset := range prepackagedPlugins {
		if tarball, err := asset("plugin.tar.gz"); err != nil {
			l4g.Error("failed to install prepackaged plugin: " + err.Error())
		} else if tarball != nil {
			a.removePlugin(id, true)
			if _, err := a.installPlugin(bytes.NewReader(tarball), true); err != nil {
				l4g.Error("failed to install prepackaged plugin: " + err.Error())
			}
			if _, ok := a.Config().PluginSettings.PluginStates[id]; !ok {
				if err := a.EnablePlugin(id); err != nil {
					l4g.Error("failed to enable prepackaged plugin: " + err.Error())
				}
			}
		}
	}

	utils.RemoveConfigListener(a.PluginConfigListenerId)
	a.PluginConfigListenerId = utils.AddConfigListener(func(prevCfg, cfg *model.Config) {
		if a.PluginEnv == nil {
			return
		}

		if *prevCfg.PluginSettings.Enable && *cfg.PluginSettings.Enable {
			a.ActivatePlugins()
		}

		for _, err := range a.PluginEnv.Hooks().OnConfigurationChange() {
			l4g.Error(err.Error())
		}
	})

	a.ActivatePlugins()
}

func (a *App) ServePluginRequest(w http.ResponseWriter, r *http.Request) {
	if a.PluginEnv == nil || !*a.Config().PluginSettings.Enable {
		err := model.NewAppError("ServePluginRequest", "app.plugin.disabled.app_error", nil, "Enable plugins to serve plugin requests", http.StatusNotImplemented)
		err.Translate(utils.T)
		l4g.Error(err.Error())
		w.WriteHeader(err.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.ToJson()))
		return
	}

	token := ""

	authHeader := r.Header.Get(model.HEADER_AUTH)
	if strings.HasPrefix(strings.ToUpper(authHeader), model.HEADER_BEARER+" ") {
		token = authHeader[len(model.HEADER_BEARER)+1:]
	} else if strings.HasPrefix(strings.ToLower(authHeader), model.HEADER_TOKEN+" ") {
		token = authHeader[len(model.HEADER_TOKEN)+1:]
	} else if cookie, _ := r.Cookie(model.SESSION_COOKIE_TOKEN); cookie != nil && (r.Method == "GET" || r.Header.Get(model.HEADER_REQUESTED_WITH) == model.HEADER_REQUESTED_WITH_XML) {
		token = cookie.Value
	} else {
		token = r.URL.Query().Get("access_token")
	}

	r.Header.Del("Mattermost-User-Id")
	if token != "" {
		if session, err := a.GetSession(token); err == nil {
			r.Header.Set("Mattermost-User-Id", session.UserId)
		}
	}

	cookies := r.Cookies()
	r.Header.Del("Cookie")
	for _, c := range cookies {
		if c.Name != model.SESSION_COOKIE_TOKEN {
			r.AddCookie(c)
		}
	}
	r.Header.Del(model.HEADER_AUTH)
	r.Header.Del("Referer")

	newQuery := r.URL.Query()
	newQuery.Del("access_token")
	r.URL.RawQuery = newQuery.Encode()

	params := mux.Vars(r)
	a.PluginEnv.Hooks().ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "plugin_id", params["plugin_id"])))
}

func (a *App) ShutDownPlugins() {
	if a.PluginEnv == nil {
		return
	}

	l4g.Info("Shutting down plugins")

	for _, err := range a.PluginEnv.Shutdown() {
		l4g.Error(err.Error())
	}
	utils.RemoveConfigListener(a.PluginConfigListenerId)
	a.PluginConfigListenerId = ""
	a.PluginEnv = nil
}

func (a *App) SetPluginKey(pluginId string, key string, value []byte) *model.AppError {
	kv := &model.PluginKeyValue{
		PluginId: pluginId,
		Key:      key,
		Value:    value,
	}

	result := <-a.Srv.Store.Plugin().SaveOrUpdate(kv)

	if result.Err != nil {
		l4g.Error(result.Err.Error())
	}

	return result.Err
}

func (a *App) GetPluginKey(pluginId string, key string) ([]byte, *model.AppError) {
	result := <-a.Srv.Store.Plugin().Get(pluginId, key)

	if result.Err != nil {
		if result.Err.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		l4g.Error(result.Err.Error())
		return nil, result.Err
	}

	kv := result.Data.(*model.PluginKeyValue)

	return kv.Value, nil
}

func (a *App) DeletePluginKey(pluginId string, key string) *model.AppError {
	result := <-a.Srv.Store.Plugin().Delete(pluginId, key)

	if result.Err != nil {
		l4g.Error(result.Err.Error())
	}

	return result.Err
}
