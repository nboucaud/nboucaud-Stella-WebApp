// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package slashcommands

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/app"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/shared/i18n"
)

const (
	AvailableRemoteActions = "invite, accept, remove, status"
)

type RemoteProvider struct {
}

const (
	CommandTriggerRemote = "remote"
)

func init() {
	app.RegisterCommandProvider(&RemoteProvider{})
}

func (rp *RemoteProvider) GetTrigger() string {
	return CommandTriggerRemote
}

func (rp *RemoteProvider) GetCommand(a *app.App, T i18n.TranslateFunc) *model.Command {

	remote := model.NewAutocompleteData(rp.GetTrigger(), "[action]", T("api.command_remote.remote_add_remove.help", map[string]interface{}{"Actions": AvailableRemoteActions}))

	invite := model.NewAutocompleteData("invite", "", T("api.command_remote.invite.help"))
	invite.AddNamedTextArgument("password", T("api.command_remote.invite_password.help"), T("api.command_remote.invite_password.hint"), "", true)
	invite.AddNamedTextArgument("name", T("api.command_remote.name.help"), T("api.command_remote.name.hint"), "", true)
	invite.AddNamedTextArgument("displayname", T("api.command_remote.displayname.help"), T("api.command_remote.displayname.hint"), "", false)

	accept := model.NewAutocompleteData("accept", "", T("api.command_remote.accept.help"))
	accept.AddNamedTextArgument("password", T("api.command_remote.invite_password.help"), T("api.command_remote.invite_password.hint"), "", true)
	accept.AddNamedTextArgument("name", T("api.command_remote.name.help"), T("api.command_remote.name.hint"), "", true)
	accept.AddNamedTextArgument("displayname", T("api.command_remote.displayname.help"), T("api.command_remote.displayname.hint"), "", false)
	accept.AddNamedTextArgument("invite", T("api.command_remote.invitation.help"), T("api.command_remote.invitation.hint"), "", true)

	remove := model.NewAutocompleteData("remove", "", T("api.command_remote.remove.help"))
	remove.AddNamedDynamicListArgument("remoteId", T("api.command_remote.remove_remote_id.help"), "builtin:remote", true)

	status := model.NewAutocompleteData("status", "", T("api.command_remote.status.help"))

	remote.AddCommand(invite)
	remote.AddCommand(accept)
	remote.AddCommand(remove)
	remote.AddCommand(status)

	return &model.Command{
		Trigger:          rp.GetTrigger(),
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_remote.desc"),
		AutoCompleteHint: T("api.command_remote.hint"),
		DisplayName:      T("api.command_remote.name"),
		AutocompleteData: remote,
	}
}

func (rp *RemoteProvider) DoCommand(a *app.App, args *model.CommandArgs, message string) *model.CommandResponse {
	if !a.HasPermissionTo(args.UserId, model.PERMISSION_MANAGE_SHARED_CHANNELS) {
		return responsef(args.T("api.command_remote.permission_required", map[string]interface{}{"Permission": "manage_shared_channels"}))
	}

	margs := parseNamedArgs(args.Command)
	action, ok := margs[ActionKey]
	if !ok {
		return responsef(args.T("api.command_remote.missing_command", map[string]interface{}{"Actions": AvailableRemoteActions}))
	}

	switch action {
	case "invite":
		return rp.doInvite(a, args, margs)
	case "accept":
		return rp.doAccept(a, args, margs)
	case "remove":
		return rp.doRemove(a, args, margs)
	case "status":
		return rp.doStatus(a, args, margs)
	}

	return responsef(args.T("api.command_remote.unknown_action", map[string]interface{}{"Action": action}))
}

func (rp *RemoteProvider) GetAutoCompleteListItems(a *app.App, commandArgs *model.CommandArgs, arg *model.AutocompleteArg, parsed, toBeParsed string) ([]model.AutocompleteListItem, error) {
	if !a.HasPermissionTo(commandArgs.UserId, model.PERMISSION_MANAGE_SHARED_CHANNELS) {
		return nil, errors.New("You require `manage_shared_channels` permission to manage remote clusters.")
	}

	if arg.Name == "remoteId" && strings.Contains(parsed, " remove ") {
		return getRemoteClusterAutocompleteListItems(a, true)
	}

	return nil, fmt.Errorf("`%s` is not a dynamic argument", arg.Name)
}

// doInvite creates and displays an encrypted invite that can be used by a remote site to establish a simple trust.
func (rp *RemoteProvider) doInvite(a *app.App, args *model.CommandArgs, margs map[string]string) *model.CommandResponse {
	password := margs["password"]
	if password == "" {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "password"}))
	}

	name := margs["name"]
	if name == "" {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "name"}))
	}

	displayname := margs["displayname"]
	if displayname == "" {
		displayname = name
	}

	url := a.GetSiteURL()
	if url == "" {
		return responsef(args.T("api.command_remote.site_url_not_set"))
	}

	rc := &model.RemoteCluster{
		Name:        name,
		DisplayName: displayname,
		Token:       model.NewId(),
		CreatorId:   args.UserId,
	}

	rcSaved, appErr := a.AddRemoteCluster(rc)
	if appErr != nil {
		return responsef(args.T("api.command_remote.add_remote.error", map[string]interface{}{"Error": appErr.Error()}))
	}

	// Display the encrypted invitation
	invite := &model.RemoteClusterInvite{
		RemoteId:     rcSaved.RemoteId,
		RemoteTeamId: args.TeamId,
		SiteURL:      url,
		Token:        rcSaved.Token,
	}
	encrypted, err := invite.Encrypt(password)
	if err != nil {
		return responsef(args.T("api.command_remote.encrypt_invitation.error", map[string]interface{}{"Error": err.Error()}))
	}
	encoded := base64.URLEncoding.EncodeToString(encrypted)

	return responsef("##### " + args.T("api.command_remote.invitation_created") + "\n" +
		args.T("api.command_remote.invite_summary", map[string]interface{}{"Command": "/remote accept", "Invitation": encoded, "SiteURL": invite.SiteURL}))
}

// doAccept accepts an invitation generated by a remote site.
func (rp *RemoteProvider) doAccept(a *app.App, args *model.CommandArgs, margs map[string]string) *model.CommandResponse {
	password := margs["password"]
	if password == "" {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "password"}))
	}

	name := margs["name"]
	if name == "" {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "name"}))
	}

	displayname := margs["displayname"]
	if displayname == "" {
		displayname = name
	}

	blob := margs["invite"]
	if blob == "" {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "invite"}))
	}

	// invite is encoded as base64 and encrypted
	decoded, err := base64.URLEncoding.DecodeString(blob)
	if err != nil {
		return responsef(args.T("api.command_remote.decode_invitation.error", map[string]interface{}{"Error": err.Error()}))
	}
	invite := &model.RemoteClusterInvite{}
	err = invite.Decrypt(decoded, password)
	if err != nil {
		return responsef(args.T("api.command_remote.incorrect_password.error", map[string]interface{}{"Error": err.Error()}))
	}

	rcs, _ := a.GetRemoteClusterService()
	if rcs == nil {
		return responsef(args.T("api.command_remote.service_not_enabled"))
	}

	url := a.GetSiteURL()
	if url == "" {
		return responsef(args.T("api.command_remote.site_url_not_set"))
	}

	rc, err := rcs.AcceptInvitation(invite, name, displayname, args.UserId, args.TeamId, url)
	if err != nil {
		return responsef(args.T("api.command_remote.accept_invitation.error", map[string]interface{}{"Error": err.Error()}))
	}

	return responsef("##### " + args.T("api.command_remote.accept_invitation", map[string]interface{}{"SiteURL": rc.SiteURL}))
}

// doRemove removes a remote cluster from the database, effectively revoking the trust relationship.
func (rp *RemoteProvider) doRemove(a *app.App, args *model.CommandArgs, margs map[string]string) *model.CommandResponse {
	id, ok := margs["remoteId"]
	if !ok {
		return responsef(args.T("api.command_remote.missing_empty", map[string]interface{}{"Arg": "remoteId"}))
	}

	deleted, err := a.DeleteRemoteCluster(id)
	if err != nil {
		responsef(args.T("api.command_remote.remove_remote.error", map[string]interface{}{"Error": err.Error()}))
	}

	result := "removed"
	if !deleted {
		result = "**NOT FOUND**"
	}
	return responsef("##### " + args.T("api.command_remote.cluster_removed", map[string]interface{}{"RemoteId": id, "Result": result}))
}

// doStatus displays connection status for all remote clusters.
func (rp *RemoteProvider) doStatus(a *app.App, args *model.CommandArgs, _ map[string]string) *model.CommandResponse {
	list, err := a.GetAllRemoteClusters(model.RemoteClusterQueryFilter{})
	if err != nil {
		responsef(args.T("api.command_remote.fetch_status.error", map[string]interface{}{"Error": err.Error()}))
	}

	if len(list) == 0 {
		return responsef("** " + args.T("api.command_remote.remotes_not_found") + " **")
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, args.T("api.command_remote.remote_table_header")+"| \n")
	fmt.Fprintf(&sb, "| ---- | -------- | ---------- | :-------------: | :----: | ---------- |\n")

	for _, rc := range list {
		accepted := ":white_check_mark:"
		if rc.SiteURL == "" {
			accepted = ":x:"
		}

		online := ":white_check_mark:"
		if !isOnline(rc.LastPingAt) {
			online = ":skull_and_crossbones:"
		}

		lastPing := formatTimestamp(model.GetTimeForMillis(rc.LastPingAt))

		fmt.Fprintf(&sb, "| %s | %s | %s | %s | %s | %s | %s |\n", rc.Name, rc.DisplayName, rc.SiteURL, rc.RemoteId, accepted, online, lastPing)
	}
	return responsef(sb.String())
}

func isOnline(lastPing int64) bool {
	return lastPing > model.GetMillis()-model.RemoteOfflineAfterMillis
}

func getRemoteClusterAutocompleteListItems(a *app.App, includeOffline bool) ([]model.AutocompleteListItem, error) {
	filter := model.RemoteClusterQueryFilter{
		ExcludeOffline: !includeOffline,
	}
	clusters, err := a.GetAllRemoteClusters(filter)
	if err != nil || len(clusters) == 0 {
		return []model.AutocompleteListItem{}, nil
	}

	list := make([]model.AutocompleteListItem, 0, len(clusters))

	for _, rc := range clusters {
		item := model.AutocompleteListItem{
			Item:     rc.RemoteId,
			HelpText: fmt.Sprintf("%s  (%s)", rc.DisplayName, rc.SiteURL)}
		list = append(list, item)
	}
	return list, nil
}

func getRemoteClusterAutocompleteListItemsNotInChannel(a *app.App, channelId string, includeOffline bool) ([]model.AutocompleteListItem, error) {
	filter := model.RemoteClusterQueryFilter{
		ExcludeOffline: !includeOffline,
		NotInChannel:   channelId,
	}
	all, err := a.GetAllRemoteClusters(filter)
	if err != nil || len(all) == 0 {
		return []model.AutocompleteListItem{}, nil
	}

	list := make([]model.AutocompleteListItem, 0, len(all))

	for _, rc := range all {
		item := model.AutocompleteListItem{
			Item:     rc.RemoteId,
			HelpText: fmt.Sprintf("%s  (%s)", rc.DisplayName, rc.SiteURL)}
		list = append(list, item)
	}
	return list, nil
}
