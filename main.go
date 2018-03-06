package main

const binary = "slackcli"
const version = "0.6.5"

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
	cli.Register(cli.CallFilesCommentsDelete, "files.comments.delete", []string{"file", "fcid"}, "Deletes an existing comment on a file")
	cli.Register(cli.CallFilesCommentsEdit, "files.comments.edit", []string{"file", "fcid", "text"}, "Edit an existing file comment")
	cli.Register(cli.CallFilesDelete, "files.delete", []string{"file"}, "Deletes a file and associated comments")
	cli.Register(cli.CallFilesInfo, "files.info", []string{"file", "count", "page"}, "Gets information about a team file")
	cli.Register(cli.CallFilesList, "files.list", []string{"count", "page"}, "Lists and filters team files")
	cli.Register(cli.CallFilesListAfterTime, "files.listAfterTime", []string{"time", "count", "page"}, "Lists and filters team files after this timestamp (inclusive)")
	cli.Register(cli.CallFilesListBeforeTime, "files.listBeforeTime", []string{"time", "count", "page"}, "Lists and filters team files before this timestamp (inclusive)")
	cli.Register(cli.CallFilesListByChannel, "files.listByChannel", []string{"channel", "count", "page"}, "Lists and filters team files in a specific channel")
	cli.Register(cli.CallFilesListByType, "files.listByType", []string{"type", "count", "page"}, "Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs")
	cli.Register(cli.CallFilesListByUser, "files.listByUser", []string{"user", "count", "page"}, "Lists and filters team files created by a single user")
	cli.Register(cli.CallFilesRevokePublicURL, "files.revokePublicURL", []string{"file"}, "Revokes public/external sharing access for a file")
	cli.Register(cli.CallFilesSharedPublicURL, "files.sharedPublicURL", []string{"file"}, "Enables a file for public/external sharing")
	cli.Register(cli.CallFilesUpload, "files.upload", []string{"channel", "filename"}, "Uploads or creates a file from local data")
	cli.Register(cli.CallGroupsArchive, "groups.archive", []string{"channel"}, "Archives a private channel")
	cli.Register(cli.CallGroupsClose, "groups.close", []string{"channel"}, "Closes a private channel")
	cli.Register(cli.CallGroupsCreate, "groups.create", []string{"channel"}, "Creates a private channel")
	cli.Register(cli.CallGroupsCreateChild, "groups.createChild", []string{"channel"}, "Clones and archives a private channel")
	cli.Register(cli.CallGroupsHistory, "groups.history", []string{"channel", "time"}, "Fetches history of messages and events from a private channel")
	cli.Register(cli.CallGroupsID, "groups.id", []string{"channel"}, "Gets private channel identifier from readable name")
	cli.Register(cli.CallGroupsInfo, "groups.info", []string{"channel"}, "Gets information about a private channel")
	cli.Register(cli.CallGroupsInvite, "groups.invite", []string{"channel", "user"}, "Invites a user to a private channel")
	cli.Register(cli.CallGroupsKick, "groups.kick", []string{"channel", "user"}, "Removes a user from a private channel")
	cli.Register(cli.CallGroupsLeave, "groups.leave", []string{"channel"}, "Leaves a private channel")
	cli.Register(cli.CallGroupsList, "groups.list", []string{}, "Lists private channels that the calling user has access to")
	cli.Register(cli.CallGroupsMark, "groups.mark", []string{"channel", "time"}, "Sets the read cursor in a private channel")
	cli.Register(cli.CallGroupsMyHistory, "groups.myHistory", []string{"channel", "time"}, "Displays messages of the current user from a private channel")
	cli.Register(cli.CallGroupsOpen, "groups.open", []string{"group"}, "Opens a private channel")
	cli.Register(cli.CallGroupsRename, "groups.rename", []string{"channel", "name"}, "Renames a private channel")
	cli.Register(cli.CallGroupsPurgeHistory, "groups.purgeHistory", []string{"channel", "time"}, "Deletes history of messages and events from a private channel")
	cli.Register(cli.CallGroupsSetPurpose, "groups.setPurpose", []string{"channel", "purpose"}, "Sets the purpose for a private channel")
	cli.Register(cli.CallGroupsSetRetention, "groups.setRetention", []string{"channel", "duration"}, "Sets the retention time of the messages")
	cli.Register(cli.CallGroupsSetTopic, "groups.setTopic", []string{"channel", "topic"}, "Sets the topic for a private channel")
	cli.Register(cli.CallGroupsUnarchive, "groups.unarchive", []string{"channel"}, "Unarchives a private channel")
	cli.Register(cli.CallHelpIssuesList, "help.issues.list", []string{}, "List issues reported by the current user")
	cli.Register(cli.CallImClose, "im.close", []string{"channel"}, "Close a direct message channel")
	cli.Register(cli.CallImHistory, "im.history", []string{"channel", "time"}, "Fetches history of messages and events from direct message channel")
	cli.Register(cli.CallImList, "im.list", []string{}, "Lists direct message channels for the calling user")
	cli.Register(cli.CallImMark, "im.mark", []string{"channel", "time"}, "Sets the read cursor in a direct message channel")
	cli.Register(cli.CallImMyHistory, "im.myHistory", []string{"channel", "time"}, "Displays messages of the current user from direct message channel")
	cli.Register(cli.CallImOpen, "im.open", []string{"user"}, "Opens a direct message channel")
	cli.Register(cli.CallImPurgeHistory, "im.purgeHistory", []string{"channel", "time"}, "Deletes history of messages and events from direct message channel")
	cli.Register(cli.CallMigrationExchange, "migration.exchange", []string{"users", "order"}, "For Enterprise Grid workspaces, map local user IDs to global user IDs")
	cli.Register(cli.CallMpimClose, "mpim.close", []string{"channel"}, "Closes a multiparty direct message channel")
	cli.Register(cli.CallMpimHistory, "mpim.history", []string{"channel", "time"}, "Fetches history of messages and events from a multiparty direct message")
	cli.Register(cli.CallMpimList, "mpim.list", []string{}, "Lists multiparty direct message channels for the calling user")
	cli.Register(cli.CallMpimListSimple, "mpim.listSimple", []string{}, "Lists ID and members in a multiparty direct message channels")
	cli.Register(cli.CallMpimMark, "mpim.mark", []string{"channel", "time"}, "Sets the read cursor in a multiparty direct message channel")
	cli.Register(cli.CallMpimMyHistory, "mpim.myHistory", []string{"channel", "time"}, "Displays messages of the current user from multiparty direct message channel")
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	cli.Execute()
}
