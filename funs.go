package main

import (
	"flag"
	"fmt"

	"github.com/cixtor/slackapi"
)

// CallHelp prints the description and usage options.
func (cli *CLI) CallHelp() int {
	fmt.Println("Slack CLI v" + version)
	fmt.Println("  http://cixtor.com/")
	fmt.Println("  https://api.slack.com/")
	fmt.Println("  https://github.com/cixtor/slackcli")
	fmt.Println("  https://github.com/cixtor/slackapi")
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

// CallVersion prints the program version.
func (cli *CLI) CallVersion() int {
	fmt.Printf("{\"version\":\"%s\"}\n", version)
	return 1
}
