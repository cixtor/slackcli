package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cixtor/slackapi"
)

// CallHelp prints the description and usage options.
func (cli *CLI) CallHelp() int {
	flag.Usage()
	return 2
}

// CallAPITest sends a http request with the api.test action.
func (cli *CLI) CallAPITest() int {
	return cli.PrintJSON(cli.api.APITest(flag.Arg(1)))
}

// CallAppsConnectionsOpen sends a http request with the apps.connections.open action.
func (cli *CLI) CallAppsConnectionsOpen() int {
	return cli.PrintJSON(cli.api.AppsConnectionsOpen())
}

// CallAppsEventAuthorizationsList sends a http request with the apps.event.authorizations.list action.
func (cli *CLI) CallAppsEventAuthorizationsList() int {
	return cli.PrintJSON(cli.api.AppsEventAuthorizationsList(slackapi.AppsEventAuthorizationsListInput{
		EventContext: flag.Arg(1),
		Cursor:       flag.Arg(2),
		Limit:        cli.Number(3, 100),
	}))
}

// CallAppsList sends a http request with the apps.list action.
func (cli *CLI) CallAppsList() int {
	return cli.PrintJSON(cli.api.AppsList())
}

// CallAppsManifestCreate sends a http request with the apps.manifest.create action.
func (cli *CLI) CallAppsManifestCreate() int {
	return cli.PrintJSON(cli.api.AppsManifestCreate(flag.Arg(1)))
}

// CallAppsManifestDelete sends a http request with the apps.manifest.delete action.
func (cli *CLI) CallAppsManifestDelete() int {
	return cli.PrintJSON(cli.api.AppsManifestDelete(flag.Arg(1)))
}

// CallAppsManifestExport sends a http request with the apps.manifest.export action.
func (cli *CLI) CallAppsManifestExport() int {
	return cli.PrintJSON(cli.api.AppsManifestExport(flag.Arg(1)))
}

// CallAppsManifestUpdate sends a http request with the apps.manifest.update action.
func (cli *CLI) CallAppsManifestUpdate() int {
	return cli.PrintJSON(cli.api.AppsManifestUpdate(flag.Arg(1), flag.Arg(2)))
}

// CallAppsManifestValidate sends a http request with the apps.manifest.validate action.
func (cli *CLI) CallAppsManifestValidate() int {
	return cli.PrintJSON(cli.api.AppsManifestValidate(flag.Arg(1), flag.Arg(2)))
}

// CallAuthRevoke sends a http request with the auth.revoke action.
func (cli *CLI) CallAuthRevoke() int {
	if flag.Arg(1) == "test" {
		return cli.PrintJSON(cli.api.AuthRevoke(true))
	} else {
		return cli.PrintJSON(cli.api.AuthRevoke(false))
	}
}

// CallAuthTeamsList sends a http request with the auth.teams.list action.
func (cli *CLI) CallAuthTeamsList() int {
	return cli.PrintJSON(cli.api.AuthTeamsList(slackapi.AuthTeamsListInput{
		Cursor:      flag.Arg(1),
		IncludeIcon: flag.Arg(2) == "true",
		Limit:       cli.Number(3, 100),
	}))
}

// CallAuthTest sends a http request with the auth.test action.
func (cli *CLI) CallAuthTest() int {
	res, err := cli.api.AuthTest()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return cli.PrintJSON(res)
}

// CallBotsInfo sends a http request with the bots.info action.
func (cli *CLI) CallBotsInfo() int {
	return cli.PrintJSON(cli.api.BotsInfo(flag.Arg(1)))
}

// CallChatDelete sends a http request with the chat.delete action.
func (cli *CLI) CallChatDelete() int {
	return cli.PrintJSON(cli.api.ChatDelete(slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Ts:      flag.Arg(2),
	}))
}

// CallChatDeleteAttachment sends a http request with the chat.delete action.
func (cli *CLI) CallChatDeleteAttachment() int {
	return cli.PrintJSON(cli.api.ChatDeleteAttachment(slackapi.ChatDeleteAttachmentInput{
		Channel:    flag.Arg(1),
		Ts:         flag.Arg(2),
		Attachment: cli.Number(3, 1),
	}))
}

// CallChatMeMessage sends a http request with the chat.meMessage action.
func (cli *CLI) CallChatMeMessage() int {
	return cli.PrintJSON(cli.api.ChatMeMessage(slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Text:    flag.Arg(2),
	}))
}

// CallChatPostAttachment sends a http request with the chat.postAttachment action.
func (cli *CLI) CallChatPostAttachment() int {
	var data slackapi.Attachment

	if err := json.Unmarshal([]byte(flag.Arg(2)), &data); err != nil {
		fmt.Printf("{\"ok\":false, \"error\":\"json.decode; %s\"}\n", err.Error())
		return 1
	}

	return cli.PrintJSON(cli.api.ChatPostMessage(slackapi.MessageArgs{
		Channel:     flag.Arg(1),
		Attachments: []slackapi.Attachment{data},
	}))
}

// CallChatPostMessage sends a http request with the chat.postMessage action.
func (cli *CLI) CallChatPostMessage() int {
	return cli.PrintJSON(cli.api.ChatPostMessage(slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Text:    flag.Arg(2),
	}))
}

// CallChatRobotMessage sends a http request with the chat.robotMessage action.
func (cli *CLI) CallChatRobotMessage() int {
	robotName := os.Getenv("SLACK_ROBOT_NAME")
	robotImage := os.Getenv("SLACK_ROBOT_IMAGE")

	data := slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Text:    flag.Arg(2),
		AsUser:  false,
	}

	if robotName == "" {
		robotName = "foobar"
	}

	if robotImage == "" {
		robotImage = ":slack:"
	}

	data.Username = robotName

	if robotImage[0] == ':' {
		data.IconEmoji = robotImage
	} else {
		data.IconURL = robotImage
	}

	return cli.PrintJSON(cli.api.ChatPostMessage(data))
}

// CallChatUpdate sends a http request with the chat.update action.
func (cli *CLI) CallChatUpdate() int {
	return cli.PrintJSON(cli.api.ChatUpdate(slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Ts:      flag.Arg(2),
		Text:    flag.Arg(3),
	}))
}

// CallClientCounts sends a http request with the client.counts action.
func (cli *CLI) CallClientCounts() int {
	return cli.PrintJSON(cli.api.ClientCounts(slackapi.ClientCountsInput{
		OrgWideAware:          true,
		ThreadCountsByChannel: true,
	}))
}

// CallClientShouldReload sends a http request with the client.shouldReload action.
func (cli *CLI) CallClientShouldReload() int {
	return cli.PrintJSON(cli.api.ClientShouldReload(slackapi.ClientShouldReloadInput{
		TeamIDs:         flag.Arg(1),
		VersionTs:       cli.Number(2, 1),
		BuildVersionTs:  cli.Number(3, 1),
		ConfigVersionTs: cli.Number(4, 1),
	}))
}

// CallConversationsAcceptSharedInvite sends a http request with the conversations.acceptSharedInvite action.
func (cli *CLI) CallConversationsAcceptSharedInvite() int {
	return cli.PrintJSON(cli.api.ConversationsAcceptSharedInvite(slackapi.ConversationsAcceptSharedInviteInput{
		ChannelName:       flag.Arg(1),
		ChannelID:         flag.Arg(2),
		FreeTrialAccepted: flag.Arg(3) == "true",
		InviteID:          flag.Arg(4),
		IsPrivate:         flag.Arg(5) == "true",
		TeamID:            flag.Arg(6),
	}))
}

// CallConversationsApproveSharedInvite sends a http request with the conversations.acceptSharedInvite action.
func (cli *CLI) CallConversationsApproveSharedInvite() int {
	return cli.PrintJSON(cli.api.ConversationsApproveSharedInvite(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsArchive sends a http request with the conversations.archive action.
func (cli *CLI) CallConversationsArchive() int {
	return cli.PrintJSON(cli.api.ConversationsArchive(flag.Arg(1)))
}

// CallConversationsClose sends a http request with the conversations.close action.
func (cli *CLI) CallConversationsClose() int {
	return cli.PrintJSON(cli.api.ConversationsClose(flag.Arg(1)))
}

// CallConversationsCreate sends a http request with the conversations.create action.
func (cli *CLI) CallConversationsCreate() int {
	return cli.PrintJSON(cli.api.ConversationsCreate(slackapi.ConversationsCreateInput{
		Name:      flag.Arg(1),
		IsPrivate: flag.Arg(2) == "true",
		TeamID:    flag.Arg(3),
	}))
}

// CallConversationsDeclineSharedInvite sends a http request with the conversations.declineSharedInvite action.
func (cli *CLI) CallConversationsDeclineSharedInvite() int {
	return cli.PrintJSON(cli.api.ConversationsDeclineSharedInvite(slackapi.ConversationsDeclineSharedInviteInput{
		InviteID:   flag.Arg(1),
		TargetTeam: flag.Arg(2),
	}))
}

// CallConversationsDelete sends a http request with the conversations.delete action.
func (cli *CLI) CallConversationsDelete() int {
	return cli.PrintJSON(cli.api.ConversationsDelete(flag.Arg(1)))
}

// CallConversationsGenericInfo sends a http request with the conversations.genericInfo action.
func (cli *CLI) CallConversationsGenericInfo() int {
	return cli.PrintJSON(cli.api.ConversationsGenericInfo(flag.Arg(1)))
}

// CallConversationsHistory sends a http request with the conversations.history action.
func (cli *CLI) CallConversationsHistory() int {
	return cli.PrintJSON(cli.api.ConversationsHistory(
		slackapi.ConversationsHistoryInput{
			Channel: flag.Arg(1),
			Latest:  flag.Arg(2),
		},
	))
}

// CallConversationsID sends a http request with the conversations.id action.
func (cli *CLI) CallConversationsID() int {
	channel := flag.Arg(1)
	result := cli.api.SearchModules(slackapi.SearchModulesInput{
		Module:            "channels",
		Query:             channel,
		Count:             cli.Number(2, 100),
		Page:              cli.Number(3, 1),
		Sort:              "timestamp",
		SortDir:           "desc",
		ExcludeMyChannels: false,
		ExtraMessageData:  false,
		Highlight:         false,
		NoUserProfile:     false,
	})

	for _, item := range result.Items {
		if item.Name == channel {
			return cli.PrintJSON(item.ID)
		}
	}

	return 0
}

// CallConversationsInfo sends a http request with the conversations.info action.
func (cli *CLI) CallConversationsInfo() int {
	return cli.PrintJSON(cli.api.ConversationsInfo(flag.Arg(1)))
}

// CallConversationsInvite sends a http request with the conversations.invite action.
func (cli *CLI) CallConversationsInvite() int {
	return cli.PrintJSON(cli.api.ConversationsInvite(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsInviteShared sends a http request with the conversations.inviteShared action.
func (cli *CLI) CallConversationsInviteShared() int {
	return cli.PrintJSON(cli.api.ConversationsInviteShared(slackapi.ConversationsInviteSharedInput{
		Channel:         flag.Arg(1),
		Emails:          strings.Split(flag.Arg(2), ","),
		ExternalLimited: flag.Arg(3) == "true",
		UserIDs:         strings.Split(flag.Arg(4), ","),
	}))
}

// CallConversationsJoin sends a http request with the conversations.join action.
func (cli *CLI) CallConversationsJoin() int {
	return cli.PrintJSON(cli.api.ConversationsJoin(flag.Arg(1)))
}

// CallConversationsKick sends a http request with the conversations.kick action.
func (cli *CLI) CallConversationsKick() int {
	return cli.PrintJSON(cli.api.ConversationsKick(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsLeave sends a http request with the conversations.leave action.
func (cli *CLI) CallConversationsLeave() int {
	return cli.PrintJSON(cli.api.ConversationsLeave(flag.Arg(1)))
}

// CallConversationsList sends a http request with the conversations.list action.
func (cli *CLI) CallConversationsList() int {
	return cli.PrintJSON(cli.api.ConversationsList(slackapi.ConversationsListInput{}))
}

// CallConversationsListConnectInvites sends a http request with the conversations.listConnectInvites action.
func (cli *CLI) CallConversationsListConnectInvites() int {
	return cli.PrintJSON(cli.api.ConversationsListConnectInvites(slackapi.ConversationsListConnectInvitesInput{
		Count:  cli.Number(1, 100),
		Cursor: flag.Arg(2),
		TeamID: flag.Arg(3),
	}))
}

// CallConversationsMark sends a http request with the conversations.mark action.
func (cli *CLI) CallConversationsMark() int {
	return cli.PrintJSON(cli.api.ConversationsMark(slackapi.ConversationsMarkInput{
		Channel:   flag.Arg(1),
		Timestamp: flag.Arg(2),
	}))
}

// CallConversationsMembers sends a http request with the conversations.members action.
func (cli *CLI) CallConversationsMembers() int {
	return cli.PrintJSON(cli.api.ConversationsMembers(slackapi.ConversationsMembersInput{
		Channel: flag.Arg(1),
		Cursor:  flag.Arg(2),
		Limit:   cli.Number(3, 100),
	}))
}

// CallConversationsOpen sends a http request with the conversations.open action.
func (cli *CLI) CallConversationsOpen() int {
	return cli.PrintJSON(cli.api.ConversationsOpen(slackapi.ConversationsOpenInput{
		Channel:         flag.Arg(1),
		PreventCreation: flag.Arg(2) == "true",
		ReturnIm:        flag.Arg(3) == "true",
		Users:           flag.Arg(4),
	}))
}

// CallConversationsRename sends a http request with the conversations.rename action.
func (cli *CLI) CallConversationsRename() int {
	return cli.PrintJSON(cli.api.ConversationsRename(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsReplies sends a http request with the conversations.replies action.
func (cli *CLI) CallConversationsReplies() int {
	return cli.PrintJSON(cli.api.ConversationsReplies(slackapi.ConversationsRepliesInput{
		Channel:   flag.Arg(1),
		Timestamp: flag.Arg(2),
		Cursor:    flag.Arg(3),
		Inclusive: flag.Arg(4) == "true",
		Latest:    flag.Arg(5),
		Limit:     cli.Number(6, 1000),
		Oldest:    flag.Arg(7),
	}))
}

// CallConversationsSetPurpose sends a http request with the conversations.setPurpose action.
func (cli *CLI) CallConversationsSetPurpose() int {
	return cli.PrintJSON(cli.api.ConversationsSetPurpose(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsSetTopic sends a http request with the conversations.setTopic action.
func (cli *CLI) CallConversationsSetTopic() int {
	return cli.PrintJSON(cli.api.ConversationsSetTopic(flag.Arg(1), flag.Arg(2)))
}

// CallConversationsSuggestions sends a http request with the conversations.suggestions action.
func (cli *CLI) CallConversationsSuggestions() int {
	return cli.PrintJSON(cli.api.ConversationsSuggestions())
}

// CallConversationsUnarchive sends a http request with the conversations.unarchive action.
func (cli *CLI) CallConversationsUnarchive() int {
	return cli.PrintJSON(cli.api.ConversationsUnarchive(flag.Arg(1)))
}

// CallDndEndDnd sends a http request with the dnd.endDnd action.
func (cli *CLI) CallDndEndDnd() int {
	return cli.PrintJSON(cli.api.DNDEndDnd())
}

// CallDndEndSnooze sends a http request with the dnd.endSnooze action.
func (cli *CLI) CallDndEndSnooze() int {
	return cli.PrintJSON(cli.api.DNDEndSnooze())
}

// CallDndInfo sends a http request with the dnd.info action.
func (cli *CLI) CallDndInfo() int {
	return cli.PrintJSON(cli.api.DNDInfo(flag.Arg(1)))
}

// CallDndSetSnooze sends a http request with the dnd.setSnooze action.
func (cli *CLI) CallDndSetSnooze() int {
	return cli.PrintJSON(cli.api.DNDSetSnooze(cli.Number(1, 60)))
}

// CallDndTeamInfo sends a http request with the dnd.teamInfo action.
func (cli *CLI) CallDndTeamInfo() int {
	return cli.PrintJSON(cli.api.DNDTeamInfo(flag.Arg(1)))
}

// CallEmojiList sends a http request with the emoji.list action.
func (cli *CLI) CallEmojiList() int {
	return cli.PrintJSON(cli.api.EmojiList())
}

// CallEventlogHistory sends a http request with the eventlog.history action.
func (cli *CLI) CallEventlogHistory() int {
	return cli.PrintJSON(cli.api.EventlogHistory(flag.Arg(1)))
}

// CallFilesCommentsAdd sends a http request with the files.comments.add action.
func (cli *CLI) CallFilesCommentsAdd() int {
	return cli.PrintJSON(cli.api.FilesCommentsAdd(flag.Arg(1), flag.Arg(2)))
}

// CallFilesCommentsDelete sends a http request with the files.comments.delete action.
func (cli *CLI) CallFilesCommentsDelete() int {
	return cli.PrintJSON(cli.api.FilesCommentsDelete(flag.Arg(1), flag.Arg(2)))
}

// CallFilesCommentsEdit sends a http request with the files.comments.edit action.
func (cli *CLI) CallFilesCommentsEdit() int {
	return cli.PrintJSON(cli.api.FilesCommentsEdit(flag.Arg(1), flag.Arg(2), flag.Arg(3)))
}

// CallFilesDelete sends a http request with the files.delete action.
func (cli *CLI) CallFilesDelete() int {
	return cli.PrintJSON(cli.api.FilesDelete(flag.Arg(1)))
}

// CallFilesInfo sends a http request with the files.info action.
func (cli *CLI) CallFilesInfo() int {
	return cli.PrintJSON(cli.api.FilesInfo(
		flag.Arg(1),
		cli.Number(2, 1000),
		cli.Number(3, 1)))
}

// CallFilesList sends a http request with the files.list action.
func (cli *CLI) CallFilesList() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		Count: cli.Number(1, 1000),
		Page:  cli.Number(2, 1),
	}))
}

// CallFilesListAfterTime sends a http request with the files.listAfterTime action.
func (cli *CLI) CallFilesListAfterTime() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		TsFrom: flag.Arg(1),
		Count:  cli.Number(2, 1000),
		Page:   cli.Number(3, 1),
	}))
}

// CallFilesListBeforeTime sends a http request with the files.listBeforeTime action.
func (cli *CLI) CallFilesListBeforeTime() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		TsTo:  flag.Arg(1),
		Count: cli.Number(2, 1000),
		Page:  cli.Number(3, 1),
	}))
}

// CallFilesListByChannel sends a http request with the files.listByChannel action.
func (cli *CLI) CallFilesListByChannel() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		Channel: flag.Arg(1),
		Count:   cli.Number(2, 1000),
		Page:    cli.Number(3, 1),
	}))
}

// CallFilesListByType sends a http request with the files.listByType action.
func (cli *CLI) CallFilesListByType() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		Types: flag.Arg(1),
		Count: cli.Number(2, 1000),
		Page:  cli.Number(3, 1),
	}))
}

// CallFilesListByUser sends a http request with the files.listByUser action.
func (cli *CLI) CallFilesListByUser() int {
	return cli.PrintJSON(cli.api.FilesList(slackapi.FileListArgs{
		User:  flag.Arg(1),
		Count: cli.Number(2, 1000),
		Page:  cli.Number(3, 1),
	}))
}

// CallFilesRevokePublicURL sends a http request with the files.revokePublicURL action.
func (cli *CLI) CallFilesRevokePublicURL() int {
	return cli.PrintJSON(cli.api.FilesRevokePublicURL(flag.Arg(1)))
}

// CallFilesSharedPublicURL sends a http request with the files.sharedPublicURL action.
func (cli *CLI) CallFilesSharedPublicURL() int {
	return cli.PrintJSON(cli.api.FilesSharedPublicURL(flag.Arg(1)))
}

// CallFilesUpload sends a http request with the files.upload action.
func (cli *CLI) CallFilesUpload() int {
	var data slackapi.FileUploadArgs

	data.Channels = flag.Arg(1)
	data.File = "@" + flag.Arg(2)

	if data.File == "@" {
		fmt.Println("{\"ok\":false, \"error\":\"no file to upload\"}")
		return 1
	}

	// grab last part of the file path.
	if strings.Contains(data.File, "/") {
		index := strings.LastIndex(data.File, "/")
		data.Filename = data.File[index+1 : len(data.File)]
	} else {
		data.Filename = data.File
	}

	data.Title = data.Filename

	// convert file name into a human title.
	if strings.Contains(data.Filename, ".") {
		index := strings.Index(data.Filename, ".")
		data.Title = data.Filename[0:index]
		data.Title = strings.Replace(data.Title, "-", "\x20", -1)
	}

	return cli.PrintJSON(cli.api.FilesUpload(data))
}

// CallHelpIssuesList sends a http request with the help.issues.list action.
func (cli *CLI) CallHelpIssuesList() int {
	return cli.PrintJSON(cli.api.HelpIssuesList())
}

// CallMigrationExchange sends a http request with the migration.exchange action.
func (cli *CLI) CallMigrationExchange() int {
	users := strings.Split(flag.Arg(1), ",")
	order := flag.Arg(2) == "true"
	return cli.PrintJSON(cli.api.MigrationExchange(users, order))
}

// CallPinsAdd sends a http request with the pins.add action.
func (cli *CLI) CallPinsAdd() int {
	return cli.PrintJSON(cli.api.PinsAdd(flag.Arg(1), flag.Arg(2)))
}

// CallPinsList sends a http request with the pins.list action.
func (cli *CLI) CallPinsList() int {
	return cli.PrintJSON(cli.api.PinsList(flag.Arg(1)))
}

// CallPinsRemove sends a http request with the pins.remove action.
func (cli *CLI) CallPinsRemove() int {
	return cli.PrintJSON(cli.api.PinsRemove(flag.Arg(1), flag.Arg(2)))
}

// CallReactionsAdd sends a http request with the reactions.add action.
func (cli *CLI) CallReactionsAdd() int {
	return cli.PrintJSON(cli.api.ReactionsAdd(slackapi.ReactionArgs{
		Channel:   flag.Arg(1),
		Timestamp: flag.Arg(2),
		Name:      flag.Arg(3),
	}))
}

// CallReactionsGet sends a http request with the reactions.get action.
func (cli *CLI) CallReactionsGet() int {
	return cli.PrintJSON(cli.api.ReactionsGet(slackapi.ReactionArgs{
		Channel:   flag.Arg(1),
		Timestamp: flag.Arg(2),
	}))
}

// CallReactionsList sends a http request with the reactions.list action.
func (cli *CLI) CallReactionsList() int {
	return cli.PrintJSON(cli.api.ReactionsList(slackapi.ReactionListArgs{
		User: flag.Arg(1),
	}))
}

// CallReactionsRemove sends a http request with the reactions.remove action.
func (cli *CLI) CallReactionsRemove() int {
	return cli.PrintJSON(cli.api.ReactionsRemove(slackapi.ReactionArgs{
		Channel:   flag.Arg(1),
		Timestamp: flag.Arg(2),
		Name:      flag.Arg(3),
	}))
}

// CallRtmEvents sends a http request with the rtm.events action.
func (cli *CLI) CallRtmEvents() int {
	rtm, err := cli.api.NewRTM(slackapi.RTMInput{})

	if err != nil {
		fmt.Printf("{\"ok\":false, \"error\":\"rtm.events; %s\"}\n", err.Error())
		return 1
	}

	go rtm.ManageEvents()

	for msg := range rtm.Events {
		switch event := msg.Data.(type) {
		case *slackapi.HelloEvent:
			fmt.Println("hello; connection established")

		case *slackapi.PresenceChangeEvent:
			fmt.Println("presence;", event.User, "=>", event.Presence)

		case *slackapi.MessageEvent:
			if event.Text == "disconnect" {
				rtm.Disconnect()
			} else {
				fmt.Printf(
					"message; %s@%s: %#v\n",
					event.User,
					event.Channel,
					event.Text,
				)
			}

		case *slackapi.ErrorEvent:
			fmt.Println("error;", event.Text)

		case *slackapi.ReconnectURLEvent:
			fmt.Println("reconnect;", event.URL)

		default:
			fmt.Printf("%s; %#v\n", msg.Type, msg.Data)
		}
	}

	fmt.Println("stopped")
	return 0
}

// CallSignupCheckEmail sends a http request with the signup.checkEmail action.
func (cli *CLI) CallSignupCheckEmail() int {
	return cli.PrintJSON(cli.api.SignupCheckEmail(flag.Arg(1)))
}

// CallSignupConfirmEmail sends a http request with the signup.confirmEmail action.
func (cli *CLI) CallSignupConfirmEmail() int {
	return cli.PrintJSON(cli.api.SignupConfirmEmail(flag.Arg(1)))
}

// CallSearchAll sends a http request with the search.all action.
func (cli *CLI) CallSearchAll() int {
	return cli.PrintJSON(cli.api.SearchAll(slackapi.SearchArgs{
		Query:   flag.Arg(1),
		Count:   cli.Number(2, 100),
		Page:    cli.Number(3, 1),
		Sort:    "timestamp",
		SortDir: "desc",
	}))
}

// CallSearchChannels sends a http request with the search.modules action.
func (cli *CLI) CallSearchChannels() int {
	return cli.PrintJSON(cli.api.SearchModules(slackapi.SearchModulesInput{
		Module:            "channels",
		Query:             flag.Arg(1),
		Count:             cli.Number(2, 100),
		Page:              cli.Number(3, 1),
		Sort:              "timestamp",
		SortDir:           "desc",
		ExcludeMyChannels: false,
		ExtraMessageData:  true,
		Highlight:         false,
		NoUserProfile:     false,
	}))
}

// CallSearchFiles sends a http request with the search.files action.
func (cli *CLI) CallSearchFiles() int {
	return cli.PrintJSON(cli.api.SearchFiles(slackapi.SearchArgs{
		Query:   flag.Arg(1),
		Count:   cli.Number(2, 100),
		Page:    cli.Number(3, 1),
		Sort:    "timestamp",
		SortDir: "desc",
	}))
}

// CallSearchMessages sends a http request with the search.messages action.
func (cli *CLI) CallSearchMessages() int {
	return cli.PrintJSON(cli.api.SearchMessages(slackapi.SearchArgs{
		Query:   flag.Arg(1),
		Count:   cli.Number(2, 100),
		Page:    cli.Number(3, 1),
		Sort:    "timestamp",
		SortDir: "desc",
	}))
}

// CallSearchModules sends a http request with the search.modules action.
func (cli *CLI) CallSearchModules() int {
	return cli.PrintJSON(cli.api.SearchModules(slackapi.SearchModulesInput{
		Module:            flag.Arg(1),
		Query:             flag.Arg(2),
		Count:             cli.Number(3, 100),
		Page:              cli.Number(4, 1),
		Sort:              "timestamp",
		SortDir:           "desc",
		ExcludeMyChannels: false,
		ExtraMessageData:  true,
		Highlight:         false,
		NoUserProfile:     false,
	}))
}

// CallSearchUsers sends a http request with the search.messages action.
func (cli *CLI) CallSearchUsers() int {
	out, err := cli.api.SearchUsers(slackapi.SearchUsersArgs{
		Query: flag.Arg(1),
		Count: cli.Number(2, 100),
	})
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return cli.PrintJSON(out)
}

// CallStarsAdd sends a http request with the stars.add action.
func (cli *CLI) CallStarsAdd() int {
	return cli.PrintJSON(cli.api.StarsAdd(flag.Arg(1), flag.Arg(2)))
}

// CallStarsList sends a http request with the stars.list action.
func (cli *CLI) CallStarsList() int {
	return cli.PrintJSON(cli.api.StarsList(cli.Number(1, 1000), cli.Number(2, 1)))
}

// CallStarsRemove sends a http request with the stars.remove action.
func (cli *CLI) CallStarsRemove() int {
	return cli.PrintJSON(cli.api.StarsRemove(flag.Arg(1), flag.Arg(2)))
}

// CallTeamAccessLogs sends a http request with the team.accessLogs action.
func (cli *CLI) CallTeamAccessLogs() int {
	return cli.PrintJSON(cli.api.TeamAccessLogs(slackapi.TeamAccessLogsInput{
		Before: flag.Arg(1),
		Count:  cli.Number(2, 1000),
		Page:   cli.Number(3, 1),
	}))
}

// CallTeamBillableInfo sends a http request with the team.billableInfo action.
func (cli *CLI) CallTeamBillableInfo() int {
	return cli.PrintJSON(cli.api.TeamBillableInfo(flag.Arg(1), flag.Arg(2)))
}

// CallTeamBillingInfo sends a http request with the team.billing.info action.
func (cli *CLI) CallTeamBillingInfo() int {
	return cli.PrintJSON(cli.api.TeamBillingInfo())
}

// CallTeamChannelsInfo sends a http request with the team.channels.info action.
func (cli *CLI) CallTeamChannelsInfo() int {
	return cli.PrintJSON(cli.api.TeamChannelsInfo(slackapi.TeamChannelsInfoInput{
		TeamID:          flag.Arg(1),
		ChannelIDs:      strings.Split(flag.Arg(2), ","),
		CheckMembership: true,
	}))
}

// CallTeamChannelsMembership sends a http request with the team.channels.membership action.
func (cli *CLI) CallTeamChannelsMembership() int {
	return cli.PrintJSON(cli.api.TeamChannelsMembership(slackapi.TeamChannelsMembershipInput{
		TeamID:  flag.Arg(1),
		Channel: flag.Arg(2),
		UserIDs: strings.Split(flag.Arg(3), ","),
	}))
}

// CallTeamInfo sends a http request with the team.info action.
func (cli *CLI) CallTeamInfo() int {
	return cli.PrintJSON(cli.api.TeamInfo(flag.Arg(1)))
}

// CallTeamIntegrationLogs sends a http request with the team.integrationLogs action.
func (cli *CLI) CallTeamIntegrationLogs() int {
	return cli.PrintJSON(cli.api.TeamIntegrationLogs(slackapi.TeamIntegrationLogsInput{
		AppID:      flag.Arg(1),
		ChangeType: flag.Arg(2),
		Count:      flag.Arg(3),
		Page:       flag.Arg(4),
		ServiceID:  flag.Arg(5),
		TeamID:     flag.Arg(6),
		User:       flag.Arg(7),
	}))
}

// CallTeamListExternal sends a http request with the team.listExternal action.
func (cli *CLI) CallTeamListExternal() int {
	return cli.PrintJSON(cli.api.TeamListExternal(slackapi.TeamListExternalInput{
		IncludeAllVisible:   1,
		IncludeApprovedOrgs: 1,
	}))
}

// CallTeamPreferencesList sends a http request with the team.preferences.list action.
func (cli *CLI) CallTeamPreferencesList() int {
	return cli.PrintJSON(cli.api.TeamPreferencesList())
}

// CallTeamProfileGet sends a http request with the team.profile.get action.
func (cli *CLI) CallTeamProfileGet() int {
	return cli.PrintJSON(cli.api.TeamProfileGet())
}

// CallUsersCounts sends a http request with the users.counts action.
func (cli *CLI) CallUsersCounts() int {
	return cli.PrintJSON(cli.api.UsersCounts())
}

// CallUsersDeletePhoto sends a http request with the users.deletePhoto action.
func (cli *CLI) CallUsersDeletePhoto() int {
	return cli.PrintJSON(cli.api.UsersDeletePhoto())
}

// CallUsersGetPresence sends a http request with the users.getPresence action.
func (cli *CLI) CallUsersGetPresence() int {
	return cli.PrintJSON(cli.api.UsersGetPresence(flag.Arg(1)))
}

// CallUsersID sends a http request with the users.id action.
func (cli *CLI) CallUsersID() int {
	fmt.Printf(
		"{\"ok\":true, \"id\":\"%s\"}\n",
		cli.api.UsersID(flag.Arg(1), cli.Number(2, 100)),
	)
	return 0
}

// CallUsersIdentity sends a http request with the users.identity action.
func (cli *CLI) CallUsersIdentity() int {
	return cli.PrintJSON(cli.api.UsersIdentity())
}

// CallUsersInfo sends a http request with the users.info action.
func (cli *CLI) CallUsersInfo() int {
	return cli.PrintJSON(cli.api.UsersInfo(flag.Arg(1)))
}

// CallUsersList sends a http request with the users.list action.
func (cli *CLI) CallUsersList() int {
	return cli.PrintJSON(cli.api.UsersList(cli.Number(1, 100), flag.Arg(2)))
}

// CallUsersLookupByEmail sends a http request with the users.lookupByEmail action.
func (cli *CLI) CallUsersLookupByEmail() int {
	return cli.PrintJSON(cli.api.UsersLookupByEmail(flag.Arg(1)))
}

// CallUsersPrefsGet sends a http request with the users.prefs.get action.
func (cli *CLI) CallUsersPrefsGet() int {
	return cli.PrintJSON(cli.api.UsersPrefsGet())
}

// CallUsersPrefsSet sends a http request with the users.prefs.set action.
func (cli *CLI) CallUsersPrefsSet() int {
	return cli.PrintJSON(cli.api.UsersPrefsSet(flag.Arg(1), flag.Arg(2)))
}

// CallUsersPreparePhoto sends a http request with the users.preparePhoto action.
func (cli *CLI) CallUsersPreparePhoto() int {
	return cli.PrintJSON(cli.api.UsersPreparePhoto(flag.Arg(1)))
}

// CallUsersProfileGet sends a http request with the users.profile.get action.
func (cli *CLI) CallUsersProfileGet() int {
	return cli.PrintJSON(cli.api.UsersProfileGet(flag.Arg(1)))
}

// CallUsersProfileSet sends a http request with the users.profile.set action.
func (cli *CLI) CallUsersProfileSet() int {
	return cli.PrintJSON(cli.api.UsersProfileSet(flag.Arg(1), flag.Arg(2)))
}

// CallUsersSetActive sends a http request with the users.setActive action.
func (cli *CLI) CallUsersSetActive() int {
	return cli.PrintJSON(cli.api.UsersSetActive())
}

// CallUsersSetAvatar sends a http request with the users.setAvatar action.
func (cli *CLI) CallUsersSetAvatar() int {
	return cli.PrintJSON(cli.api.UsersSetAvatar(flag.Arg(1)))
}

// CallUsersSetEmail sends a http request with the users.setEmail action.
func (cli *CLI) CallUsersSetEmail() int {
	return cli.PrintJSON(cli.api.UsersProfileSet("email", flag.Arg(1)))
}

// CallUsersSetPhoto sends a http request with the users.setPhoto action.
func (cli *CLI) CallUsersSetPhoto() int {
	return cli.PrintJSON(cli.api.UsersSetPhoto(flag.Arg(1)))
}

// CallUsersSetPresence sends a http request with the users.setPresence action.
func (cli *CLI) CallUsersSetPresence() int {
	return cli.PrintJSON(cli.api.UsersSetPresence(flag.Arg(1)))
}

// CallUsersSetStatus sends a http request with the users.setStatus action.
func (cli *CLI) CallUsersSetStatus() int {
	return cli.PrintJSON(cli.api.UsersSetStatus(flag.Arg(1), flag.Arg(2)))
}

// CallUsersSetUsername sends a http request with the users.setUsername action.
func (cli *CLI) CallUsersSetUsername() int {
	return cli.PrintJSON(cli.api.UsersProfileSet("username", flag.Arg(1)))
}

// CallWorkflowsStepCompleted sends a http request with the workflows.stepCompleted action.
func (cli *CLI) CallWorkflowsStepCompleted() int {
	return cli.PrintJSON(cli.api.WorkflowsStepCompleted(slackapi.WorkflowsStepCompletedInput{
		WorkflowStepExecuteID: flag.Arg(1),
	}))
}

// CallWorkflowsStepFailed sends a http request with the workflows.stepFailed action.
func (cli *CLI) CallWorkflowsStepFailed() int {
	return cli.PrintJSON(cli.api.WorkflowsStepFailed(slackapi.WorkflowsStepFailedInput{
		WorkflowStepExecuteID: flag.Arg(1),
		Error:                 slackapi.WorkflowError{Message: flag.Arg(2)},
	}))
}

// CallWorkflowsUpdateStep sends a http request with the workflows.updateStep action.
func (cli *CLI) CallWorkflowsUpdateStep() int {
	return cli.PrintJSON(cli.api.WorkflowsUpdateStep(slackapi.WorkflowsUpdateStepInput{
		WorkflowStepExecuteID: flag.Arg(1),
		StepImageURL:          flag.Arg(2),
		StepName:              flag.Arg(3),
	}))
}
