package main

const binary = "slackcli"
const version = "0.0.12"

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
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	cli.Execute()
}
