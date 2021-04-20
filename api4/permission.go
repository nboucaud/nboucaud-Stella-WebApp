package api4

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (api *API) InitPermissions() {
	api.BaseRoutes.Permissions.Handle("/ancillary", api.ApiSessionRequired(appendAncillaryPermissions)).Methods("GET")
}

func appendAncillaryPermissions(c *Context, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["subsection_permissions"]

	if !ok || len(keys[0]) < 1 {
		c.SetInvalidUrlParam("subsection_permissions")
		return
	}

	permissions := strings.Split(keys[0], ",")
	b, err := json.Marshal(model.AddAncillaryPermissions(permissions))
	if err != nil {
		c.SetJSONEncodingError()
		return
	}
	w.Write(b)
}
