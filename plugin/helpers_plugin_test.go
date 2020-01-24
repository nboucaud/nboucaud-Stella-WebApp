// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package plugin_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest/mock"
	"github.com/mattermost/mattermost-server/v5/utils/fileutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstallPluginFromURL(t *testing.T) {
	replace := true

	t.Run("incompatible server version", func(t *testing.T) {
		h := &plugin.HelpersImpl{}
		api := &plugintest.API{}
		api.On("GetServerVersion").Return("5.1.0")
		h.API = api

		_, err := h.InstallPluginFromURL("", true)

		assert.Error(t, err)
		assert.Equal(t, "incompatible server version for plugin, minimum required version: 5.18.0, current version: 5.1.0", err.Error())
	})

	t.Run("error while parsing the download url", func(t *testing.T) {
		h := &plugin.HelpersImpl{}
		api := &plugintest.API{}
		api.On("GetServerVersion").Return("5.19.0")
		h.API = api
		_, err := h.InstallPluginFromURL("http://%41:8080/", replace)

		assert.Error(t, err)
		assert.Equal(t, "error while parsing url: parse http://%41:8080/: invalid URL escape \"%41\"", err.Error())
	})

	t.Run("errors out while downloading file", func(t *testing.T) {
		h := &plugin.HelpersImpl{}
		api := &plugintest.API{}
		api.On("GetServerVersion").Return("5.19.0")
		h.API = api
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusInternalServerError)
		}))
		defer testServer.Close()
		url := testServer.URL

		_, err := h.InstallPluginFromURL(url, replace)

		assert.Error(t, err)
		assert.Equal(t, "received 500 status code while downloading plugin from server", err.Error())
	})

	t.Run("downloads the file successfully", func(t *testing.T) {
		h := &plugin.HelpersImpl{}
		api := &plugintest.API{}
		api.On("GetServerVersion").Return("5.19.0")
		h.API = api
		path, _ := fileutils.FindDir("tests")
		tarData, err := ioutil.ReadFile(filepath.Join(path, "testplugin.tar.gz"))
		require.NoError(t, err)
		expectedManifest := &model.Manifest{Id: "testplugin"}
		api.On("InstallPlugin", mock.Anything, false).Return(expectedManifest, nil)

		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
			_, _ = res.Write(tarData)
		}))
		defer testServer.Close()
		url := testServer.URL

		manifest, err := h.InstallPluginFromURL(url, false)

		assert.NoError(t, err)
		assert.Equal(t, "testplugin", manifest.Id)
	})

	t.Run("the url pointing to server is incorrect", func(t *testing.T) {
		h := &plugin.HelpersImpl{}
		api := &plugintest.API{}
		api.On("GetServerVersion").Return("5.19.0")
		h.API = api
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusNotFound)
		}))
		defer testServer.Close()
		url := testServer.URL

		_, err := h.InstallPluginFromURL(url, false)

		assert.Error(t, err)
		assert.Equal(t, "received 404 status code while downloading plugin from server", err.Error())
	})
}

func TestGetPluginAssetURL(t *testing.T) {
	siteURL := "https://mattermost.example.com"
	api := &plugintest.API{}
	api.On("GetConfig").Return(&model.Config{ServiceSettings: model.ServiceSettings{SiteURL: &siteURL}})

	p := &plugin.HelpersImpl{API: api}

	t.Run("Valid asset directory was provided", func(t *testing.T) {
		pluginID := "mattermost-1234"
		dir := "assets"
		wantedURL := &url.URL{Scheme: "https", Host: "mattermost.example.com", Path: "/mattermost-1234/assets"}
		gotURL, err := p.GetPluginAssetURL(pluginID, dir)

		assert.Equalf(t, wantedURL, gotURL, "GetPluginAssetURL(%q, %q) got=%q; want=%v", pluginID, dir, gotURL, wantedURL)
		assert.NoError(t, err)
	})

	t.Run("Valid pluginID was provided", func(t *testing.T) {
		pluginID := "mattermost-1234"
		dir := "assets"
		wantedURL := &url.URL{Scheme: "https", Host: "mattermost.example.com", Path: "/mattermost-1234/assets"}
		gotURL, err := p.GetPluginAssetURL(pluginID, dir)

		assert.Equalf(t, wantedURL, gotURL, "GetPluginAssetURL(%q, %q) got=%q; want=%v", pluginID, dir, gotURL, wantedURL)
		assert.NoError(t, err)
	})

	t.Run("Invalid asset directory name was provided", func(t *testing.T) {
		pluginID := "mattermost-1234"
		dir := ""
		gotURL, err := p.GetPluginAssetURL(pluginID, dir)

		assert.Nilf(t, gotURL, "GetPluginAssetURL(%q, %q) got=%s; want=nil", pluginID, dir, gotURL)
		assert.Error(t, err)
	})

	t.Run("Invalid pluginID was provided", func(t *testing.T) {
		pluginID := ""
		dir := "assets"
		gotURL, err := p.GetPluginAssetURL(pluginID, dir)

		assert.Nilf(t, gotURL, "GetPluginAssetURL(%q, %q) got=%q; want=nil", pluginID, dir, gotURL)
		assert.Error(t, err)
	})

	siteURL = ""
	api.On("GetConfig").Return(&model.Config{ServiceSettings: model.ServiceSettings{SiteURL: &siteURL}})

	t.Run("Empty SiteURL was configured", func(t *testing.T) {
		pluginID := "mattermost-1234"
		dir := "assets"
		gotURL, err := p.GetPluginAssetURL(pluginID, dir)

		assert.Nilf(t, gotURL, "GetPluginAssetURL(%q, %q) got=%q; want=nil", pluginID, dir, gotURL)
		assert.Error(t, err)
	})
}
