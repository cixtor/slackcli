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
	fmt.Println("Slack CLI v" + version)
	fmt.Println("  http://cixtor.com/")
	fmt.Println("  https://api.slack.com/")
	fmt.Println("  https://github.com/cixtor/slackapi")
	fmt.Println("  https://github.com/cixtor/slackcli")
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println("  Low level Slack API client with custom commands. Slack, the 'messaging app for")
	fmt.Println("  teams' offers an API that has been used to build multiple projects around it,")
	fmt.Println("  from bots to independent clients as well as integrations with other external")
	fmt.Println("  services. This project aims to offer a low level experience for advanced users")
	fmt.Println("  that want to either drop the web client or interact with the API for testing")
	fmt.Println("  purpose.")
	fmt.Println()
	fmt.Println("Usage:")

	cli.PrintCommands()

	return 0
}

// CallAPITest sends a http request with the api.test action.
func (cli *CLI) CallAPITest() int {
	return cli.PrintJSON(cli.api.APITest())
}

// CallAppsList sends a http request with the apps.list action.
func (cli *CLI) CallAppsList() int {
	return cli.PrintJSON(cli.api.AppsList())
}

// CallAuthRevoke sends a http request with the auth.revoke action.
func (cli *CLI) CallAuthRevoke() int {
	return cli.PrintJSON(cli.api.AuthRevoke())
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

// CallChannelsHistory sends a http request with the channels.history action.
func (cli *CLI) CallChannelsHistory() int {
	return cli.PrintJSON(cli.api.ChannelsHistory(slackapi.HistoryArgs{
		Channel: flag.Arg(1),
		Latest:  flag.Arg(2),
	}))
}

// CallChannelsInfo sends a http request with the channels.info action.
func (cli *CLI) CallChannelsInfo() int {
	return cli.PrintJSON(cli.api.ChannelsInfo(flag.Arg(1)))
}

// CallChannelsInvite sends a http request with the channels.invite action.
func (cli *CLI) CallChannelsInvite() int {
	return cli.PrintJSON(cli.api.ChannelsInvite(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsJoin sends a http request with the channels.join action.
func (cli *CLI) CallChannelsJoin() int {
	return cli.PrintJSON(cli.api.ChannelsJoin(flag.Arg(1)))
}

// CallChannelsKick sends a http request with the channels.kick action.
func (cli *CLI) CallChannelsKick() int {
	return cli.PrintJSON(cli.api.ChannelsKick(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsLeave sends a http request with the channels.leave action.
func (cli *CLI) CallChannelsLeave() int {
	return cli.PrintJSON(cli.api.ChannelsLeave(flag.Arg(1)))
}

// CallChannelsList sends a http request with the channels.list action.
func (cli *CLI) CallChannelsList() int {
	return cli.PrintJSON(cli.api.ChannelsList())
}

// CallChannelsMark sends a http request with the channels.mark action.
func (cli *CLI) CallChannelsMark() int {
	return cli.PrintJSON(cli.api.ChannelsMark(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsMyHistory sends a http request with the channels.myHistory action.
func (cli *CLI) CallChannelsMyHistory() int {
	return cli.PrintJSON(cli.api.ChannelsMyHistory(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsPurgeHistory sends a http request with the channels.purgeHistory action.
func (cli *CLI) CallChannelsPurgeHistory() int {
	cli.api.ChannelsPurgeHistory(flag.Arg(1), flag.Arg(2), true)
	return 0
}

// CallChannelsRename sends a http request with the channels.rename action.
func (cli *CLI) CallChannelsRename() int {
	return cli.PrintJSON(cli.api.ChannelsRename(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsSetPurpose sends a http request with the channels.setPurpose action.
func (cli *CLI) CallChannelsSetPurpose() int {
	return cli.PrintJSON(cli.api.ChannelsSetPurpose(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsSetRetention sends a http request with the channels.setRetention action.
func (cli *CLI) CallChannelsSetRetention() int {
	return cli.PrintJSON(cli.api.ChannelsSetRetention(flag.Arg(1), cli.Number(2, 0)))
}

// CallChannelsSetTopic sends a http request with the channels.setTopic action.
func (cli *CLI) CallChannelsSetTopic() int {
	return cli.PrintJSON(cli.api.ChannelsSetTopic(flag.Arg(1), flag.Arg(2)))
}

// CallChannelsSuggestions sends a http request with the channels.suggestions action.
func (cli *CLI) CallChannelsSuggestions() int {
	return cli.PrintJSON(cli.api.ChannelsSuggestions())
}

// CallChannelsUnarchive sends a http request with the channels.unarchive action.
func (cli *CLI) CallChannelsUnarchive() int {
	return cli.PrintJSON(cli.api.ChannelsUnarchive(flag.Arg(1)))
}

// CallChatDelete sends a http request with the chat.delete action.
func (cli *CLI) CallChatDelete() int {
	return cli.PrintJSON(cli.api.ChatDelete(slackapi.MessageArgs{
		Channel: flag.Arg(1),
		Ts:      flag.Arg(2),
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

// CallGroupsClose sends a http request with the groups.close action.
func (cli *CLI) CallGroupsClose() int {
	return cli.PrintJSON(cli.api.GroupsClose(flag.Arg(1)))
}

// CallGroupsCreateChild sends a http request with the groups.createChild action.
func (cli *CLI) CallGroupsCreateChild() int {
	return cli.PrintJSON(cli.api.GroupsCreateChild(flag.Arg(1)))
}

// CallGroupsHistory sends a http request with the groups.history action.
func (cli *CLI) CallGroupsHistory() int {
	return cli.PrintJSON(cli.api.GroupsHistory(slackapi.HistoryArgs{
		Channel: flag.Arg(1),
		Latest:  flag.Arg(2),
	}))
}

// CallGroupsInfo sends a http request with the groups.info action.
func (cli *CLI) CallGroupsInfo() int {
	return cli.PrintJSON(cli.api.GroupsInfo(flag.Arg(1)))
}

// CallGroupsInvite sends a http request with the groups.invite action.
func (cli *CLI) CallGroupsInvite() int {
	return cli.PrintJSON(cli.api.GroupsInvite(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsKick sends a http request with the groups.kick action.
func (cli *CLI) CallGroupsKick() int {
	return cli.PrintJSON(cli.api.GroupsKick(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsLeave sends a http request with the groups.leave action.
func (cli *CLI) CallGroupsLeave() int {
	return cli.PrintJSON(cli.api.GroupsLeave(flag.Arg(1)))
}

// CallGroupsList sends a http request with the groups.list action.
func (cli *CLI) CallGroupsList() int {
	return cli.PrintJSON(cli.api.GroupsList())
}

// CallGroupsMark sends a http request with the groups.mark action.
func (cli *CLI) CallGroupsMark() int {
	return cli.PrintJSON(cli.api.GroupsMark(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsMyHistory sends a http request with the groups.myHistory action.
func (cli *CLI) CallGroupsMyHistory() int {
	return cli.PrintJSON(cli.api.GroupsMyHistory(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsOpen sends a http request with the groups.open action.
func (cli *CLI) CallGroupsOpen() int {
	return cli.PrintJSON(cli.api.GroupsOpen(flag.Arg(1)))
}

// CallGroupsRename sends a http request with the groups.rename action.
func (cli *CLI) CallGroupsRename() int {
	return cli.PrintJSON(cli.api.GroupsRename(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsPurgeHistory sends a http request with the groups.purgeHistory action.
func (cli *CLI) CallGroupsPurgeHistory() int {
	cli.api.GroupsPurgeHistory(flag.Arg(1), flag.Arg(2), true)
	return 0
}

// CallGroupsSetPurpose sends a http request with the groups.setPurpose action.
func (cli *CLI) CallGroupsSetPurpose() int {
	return cli.PrintJSON(cli.api.GroupsSetPurpose(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsSetRetention sends a http request with the groups.setRetention action.
func (cli *CLI) CallGroupsSetRetention() int {
	return cli.PrintJSON(cli.api.GroupsSetRetention(flag.Arg(1), cli.Number(2, 0)))
}

// CallGroupsSetTopic sends a http request with the groups.setTopic action.
func (cli *CLI) CallGroupsSetTopic() int {
	return cli.PrintJSON(cli.api.GroupsSetTopic(flag.Arg(1), flag.Arg(2)))
}

// CallGroupsUnarchive sends a http request with the groups.unarchive action.
func (cli *CLI) CallGroupsUnarchive() int {
	return cli.PrintJSON(cli.api.GroupsUnarchive(flag.Arg(1)))
}

// CallHelpIssuesList sends a http request with the help.issues.list action.
func (cli *CLI) CallHelpIssuesList() int {
	return cli.PrintJSON(cli.api.HelpIssuesList())
}

// CallImClose sends a http request with the im.close action.
func (cli *CLI) CallImClose() int {
	return cli.PrintJSON(cli.api.InstantMessageClose(flag.Arg(1)))
}

// CallImHistory sends a http request with the im.history action.
func (cli *CLI) CallImHistory() int {
	return cli.PrintJSON(cli.api.InstantMessageHistory(slackapi.HistoryArgs{
		Channel: flag.Arg(1),
		Latest:  flag.Arg(2),
	}))
}

// CallImList sends a http request with the im.list action.
func (cli *CLI) CallImList() int {
	return cli.PrintJSON(cli.api.InstantMessageList())
}

// CallImMark sends a http request with the im.mark action.
func (cli *CLI) CallImMark() int {
	return cli.PrintJSON(cli.api.InstantMessageMark(flag.Arg(1), flag.Arg(2)))
}

// CallImMyHistory sends a http request with the im.myHistory action.
func (cli *CLI) CallImMyHistory() int {
	return cli.PrintJSON(cli.api.InstantMessageMyHistory(flag.Arg(1), flag.Arg(2)))
}

// CallImOpen sends a http request with the im.open action.
func (cli *CLI) CallImOpen() int {
	return cli.PrintJSON(cli.api.InstantMessageOpen(flag.Arg(1)))
}

// CallImPurgeHistory sends a http request with the im.purgeHistory action.
func (cli *CLI) CallImPurgeHistory() int {
	cli.api.InstantMessagePurgeHistory(flag.Arg(1), flag.Arg(2), true)
	return 0
}

// CallMigrationExchange sends a http request with the migration.exchange action.
func (cli *CLI) CallMigrationExchange() int {
	users := strings.Split(flag.Arg(1), ",")
	order := flag.Arg(2) == "true"
	return cli.PrintJSON(cli.api.MigrationExchange(users, order))
}

// CallMpimClose sends a http request with the mpim.close action.
func (cli *CLI) CallMpimClose() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageClose(flag.Arg(1)))
}

// CallMpimHistory sends a http request with the mpim.history action.
func (cli *CLI) CallMpimHistory() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageHistory(slackapi.HistoryArgs{
		Channel: flag.Arg(1),
		Latest:  flag.Arg(2),
	}))
}

// CallMpimList sends a http request with the mpim.list action.
func (cli *CLI) CallMpimList() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageList())
}

// CallMpimListSimple sends a http request with the mpim.listSimple action.
func (cli *CLI) CallMpimListSimple() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageListSimple())
}

// CallMpimMark sends a http request with the mpim.mark action.
func (cli *CLI) CallMpimMark() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageMark(flag.Arg(1), flag.Arg(2)))
}

// CallMpimMyHistory sends a http request with the mpim.myHistory action.
func (cli *CLI) CallMpimMyHistory() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageMyHistory(flag.Arg(1), flag.Arg(2)))
}

// CallMpimOpen sends a http request with the mpim.open action.
func (cli *CLI) CallMpimOpen() int {
	return cli.PrintJSON(cli.api.MultiPartyInstantMessageOpen(strings.Split(flag.Arg(1), ",")))
}

// CallMpimPurgeHistory sends a http request with the mpim.purgeHistory action.
func (cli *CLI) CallMpimPurgeHistory() int {
	cli.api.MultiPartyInstantMessagePurgeHistory(flag.Arg(1), flag.Arg(2), true)
	return 0
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
	rtm, err := cli.api.NewRTM(slackapi.RTMArgs{})

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
					event.Text)
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
	return cli.PrintJSON(cli.api.TeamAccessLogs(slackapi.AccessLogArgs{
		Count: cli.Number(2, 1000),
		Page:  cli.Number(3, 1),
	}))
}

// CallTeamBillableInfo sends a http request with the team.billableInfo action.
func (cli *CLI) CallTeamBillableInfo() int {
	return cli.PrintJSON(cli.api.TeamBillableInfo(flag.Arg(1)))
}

// CallTeamInfo sends a http request with the team.info action.
func (cli *CLI) CallTeamInfo() int {
	return cli.PrintJSON(cli.api.TeamInfo())
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

// CallVersion prints the program version.
func (cli *CLI) CallVersion() int {
	fmt.Printf("{\"version\":\"%s\"}\n", version)
	return 1
}
