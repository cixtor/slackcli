package main

const binary = "slackcli"
const version = "0.3.0"

func main() {
	cli := NewCLI(binary)

	cli.AutoAuthenticate()

	cli.Register(cli.CallAPITest, "api.test", []string{}, "Checks API calling code")
	cli.Register(cli.CallAppsList, "apps.list", []string{}, "Lists associated applications")
	cli.Register(cli.CallAuthRevoke, "auth.revoke", []string{}, "Revokes a token")
	cli.Register(cli.CallAuthTest, "auth.test", []string{}, "Checks authentication and identity")
	cli.Register(cli.CallBotsInfo, "bots.info", []string{"bot"}, "Gets information about a bot user")
	cli.Register(cli.CallChannelsArchive, "channels.archive", []string{"channel"}, "Archives a channel")
	cli.Register(cli.CallChannelsCreate, "channels.create", []string{"channel"}, "Creates a channel if authorized")
	cli.Register(cli.CallChannelsHistory, "channels.history", []string{"channel", "time"}, "Fetches history of messages and events from a channel")
	cli.Register(cli.CallChannelsID, "channels.id", []string{"channel"}, "Gets channel identifier from readable name")
	cli.Register(cli.CallChannelsInfo, "channels.info", []string{"channel"}, "Gets information about a channel")
	cli.Register(cli.CallChannelsInvite, "channels.invite", []string{"channel", "user"}, "Invites a user to a channel")
	cli.Register(cli.CallChannelsJoin, "channels.join", []string{"channel"}, "Joins a channel, creates it if it does not exist")
	cli.Register(cli.CallChannelsKick, "channels.kick", []string{"channel", "user"}, "Removes a user from a channel")
	cli.Register(cli.CallChannelsLeave, "channels.leave", []string{"channel"}, "Leaves a channel")
	cli.Register(cli.CallChannelsList, "channels.list", []string{}, "Lists all channels in a Slack team")
	cli.Register(cli.CallChannelsMark, "channels.mark", []string{"channel", "time"}, "Sets the read cursor in a channel")
	cli.Register(cli.CallChannelsMyHistory, "channels.myHistory", []string{"channel", "time"}, "Displays messages of the current user from a channel")
	cli.Register(cli.CallChannelsPurgeHistory, "channels.purgeHistory", []string{"channel", "time"}, "Deletes history of messages and events from a channel")
	cli.Register(cli.CallChannelsRename, "channels.rename", []string{"channel", "name"}, "Renames a channel")
	cli.Register(cli.CallChannelsSetPurpose, "channels.setPurpose", []string{"channel", "purpose"}, "Sets the purpose for a channel")
	cli.Register(cli.CallChannelsSetRetention, "channels.setRetention", []string{"channel", "duration"}, "Sets the retention time of the messages")
	cli.Register(cli.CallChannelsSetTopic, "channels.setTopic", []string{"channel", "topic"}, "Sets the topic for a channel")
	cli.Register(cli.CallChannelsSuggestions, "channels.suggestions", []string{}, "Prints a list of suggested channels to join")
	cli.Register(cli.CallChannelsUnarchive, "channels.unarchive", []string{"channel"}, "Unarchives a channel")
	cli.Register(cli.CallChatDelete, "chat.delete", []string{"channel", "time"}, "Deletes a message")
	cli.Register(cli.CallChatMeMessage, "chat.meMessage", []string{"channel", "text"}, "Share a me message into a channel")
	cli.Register(cli.CallChatPostAttachment, "chat.postAttachment", []string{"channel", "json"}, "Sends an attachment to a channel")
	cli.Register(cli.CallChatPostMessage, "chat.postMessage", []string{"channel", "text"}, "Sends a message to a channel")
	cli.Register(cli.CallChatRobotMessage, "chat.robotMessage", []string{"channel", "text"}, "Sends a message to a channel as a robot")
	cli.Register(cli.CallChatUpdate, "chat.update", []string{"channel", "time", "text"}, "Updates a message")
	cli.Register(cli.CallDndEndDnd, "dnd.endDnd", []string{}, "Ends the current user's \"Do Not Disturb\" session immediately")
	cli.Register(cli.CallDndEndSnooze, "dnd.endSnooze", []string{}, "Ends the current user's snooze mode immediately")
	cli.Register(cli.CallDndInfo, "dnd.info", []string{"user"}, "Retrieves a user's current \"Do Not Disturb\" status")
	cli.Register(cli.CallDndSetSnooze, "dnd.setSnooze", []string{}, "Ends the current user's snooze mode immediately")
	cli.Register(cli.CallDndTeamInfo, "dnd.teamInfo", []string{"users"}, "Retrieves the \"Do Not Disturb\" status for users on a team")
	cli.Register(cli.CallEmojiList, "emoji.list", []string{}, "Lists custom emoji for a team")
	cli.Register(cli.CallEventlogHistory, "eventlog.history", []string{"time"}, "Lists all the events since the specified time")
	cli.Register(cli.CallFilesCommentsAdd, "files.comments.add", []string{"file", "text"}, "Add a comment to an existing file")
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	cli.Execute()
}
