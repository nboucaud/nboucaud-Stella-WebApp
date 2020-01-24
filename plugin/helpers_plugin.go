// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package plugin

import (
	"net/http"
	"net/url"
	"time"

	"github.com/blang/semver"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

// InstallPluginFromURL implements Helpers.InstallPluginFromURL.
func (p *HelpersImpl) InstallPluginFromURL(downloadURL string, replace bool) (*model.Manifest, error) {
	err := p.ensureServerVersion("5.18.0")
	if err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(downloadURL)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing url")
	}

	client := &http.Client{Timeout: time.Hour}
	response, err := client.Get(parsedURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "unable to download the plugin")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("received %d status code while downloading plugin from server", response.StatusCode)
	}

	manifest, installError := p.API.InstallPlugin(response.Body, replace)
	if installError != nil {
		return nil, errors.Wrap(err, "unable to install plugin on server")
	}

	return manifest, nil
}

func (p *HelpersImpl) ensureServerVersion(required string) error {
	serverVersion := p.API.GetServerVersion()
	currentVersion := semver.MustParse(serverVersion)
	requiredVersion := semver.MustParse(required)

	if currentVersion.LT(requiredVersion) {
		return errors.Errorf("incompatible server version for plugin, minimum required version: %s, current version: %s", required, serverVersion)
	}
	return nil
}

// GetPluginAssetURL builds a URL to the given asset in the assets directory.
func (p *HelpersImpl) GetPluginAssetURL(pluginID, asset string) (*url.URL, error) {
	if len(pluginID) == 0 {
		return nil, errors.New("empty pluginID provided")
	}

	if len(asset) == 0 {
		return nil, errors.New("empty asset name provided")
	}

	siteURL := *p.API.GetConfig().ServiceSettings.SiteURL
	if siteURL == "" {
		return nil, errors.New("no SiteURL configured by the server")
	}
	u, err := url.Parse(siteURL + "/" + pluginID + "/" + asset)
	if err != nil {
		return nil, err
	}
	return u, nil
}
