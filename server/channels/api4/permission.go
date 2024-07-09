// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
)

func (api *API) InitPermissions() {
	api.BaseRoutes.Permissions.Handle("/ancillary", api.APISessionRequired(appendAncillaryPermissionsPost)).Methods("POST")
}

func appendAncillaryPermissionsPost(c *Context, w http.ResponseWriter, r *http.Request) {
	permissions, err := model.NonSortedArrayFromJSON(r.Body)
	if err != nil || len(permissions) < 1 {
		c.Err = model.NewAppError("appendAncillaryPermissionsPost", model.PayloadParseError, nil, "", http.StatusBadRequest).Wrap(err)
		return
	}
	b, err := json.Marshal(model.AddAncillaryPermissions(permissions))
	if err != nil {
		c.SetJSONEncodingError(err)
		return
	}
	w.Write(b)
}
