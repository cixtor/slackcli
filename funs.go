package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func (cli *CLI) CallApiTest() int {
	return 1
}

func (cli *CLI) CallAuthTest() int {
	json.NewEncoder(os.Stdout).Encode(cli.api.AuthTest())
	return 1
}

func (cli *CLI) CallVersion() int {
	fmt.Printf("{\"version\":\"%s\"}", version)
	return 1
}
