package main

const binary = "slackcli"
const version = "0.0.20"

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
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	cli.Execute()
}
