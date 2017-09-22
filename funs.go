package main

import (
	"fmt"
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

// CallVersion prints the program version.
func (cli *CLI) CallVersion() int {
	fmt.Printf("{\"version\":\"%s\"}\n", version)
	return 1
}
