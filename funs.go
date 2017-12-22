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
	return cli.PrintJSON(cli.api.AuthTest())
}

// CallBotsInfo sends a http request with the bots.info action.
func (cli *CLI) CallBotsInfo() int {
	return cli.PrintJSON(cli.api.BotsInfo(flag.Arg(1)))
}

// CallChannelsArchive sends a http request with the channels.archive action.
func (cli *CLI) CallChannelsArchive() int {
	return cli.PrintJSON(cli.api.ChannelsArchive(flag.Arg(1)))
}

// CallChannelsCreate sends a http request with the channels.create action.
func (cli *CLI) CallChannelsCreate() int {
	return cli.PrintJSON(cli.api.ChannelsCreate(flag.Arg(1)))
}

// CallChannelsHistory sends a http request with the channels.history action.
func (cli *CLI) CallChannelsHistory() int {
	return cli.PrintJSON(cli.api.ChannelsHistory(slackapi.HistoryArgs{
		Channel: flag.Arg(1),
		Latest:  flag.Arg(2),
	}))
}

// CallChannelsID sends a http request with the channels.id action.
func (cli *CLI) CallChannelsID() int {
	fmt.Printf(
		"{\"ok\":true, \"channel\":\"%s\"}\n",
		cli.api.ChannelsID(flag.Arg(1)))
	return 0
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

	// grab last part of the file path.
	if strings.Contains(data.File, "/") {
		index := strings.LastIndex(data.File, "/")
		data.Filename = data.File[index+1 : len(data.File)]
	} else {
		data.Filename = data.File
	}

	// convert file name into a human title.
	index := strings.Index(data.Filename, ".")
	data.Title = data.Filename[0:index]
	data.Title = strings.Replace(data.Title, "-", "\x20", -1)

	return cli.PrintJSON(cli.api.FilesUpload(data))
}

// CallGroupsArchive sends a http request with the groups.archive action.
func (cli *CLI) CallGroupsArchive() int {
	return cli.PrintJSON(cli.api.GroupsArchive(flag.Arg(1)))
}

// CallGroupsClose sends a http request with the groups.close action.
func (cli *CLI) CallGroupsClose() int {
	return cli.PrintJSON(cli.api.GroupsClose(flag.Arg(1)))
}

// CallVersion prints the program version.
func (cli *CLI) CallVersion() int {
	fmt.Printf("{\"version\":\"%s\"}\n", version)
	return 1
}
