package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/cixtor/slackapi"
)

type Function func() int

type CLI struct {
	binary   string
	api      *slackapi.SlackAPI
	commands []Command
}

type Command struct {
	Function Function
	Name     string
	Params   []string
	Help     string
}

func NewCLI(binary string) *CLI {
	cli := new(CLI)

	cli.binary = binary
	cli.api = slackapi.New()

	return cli
}

func (cli *CLI) Register(fun func() int, name string, params []string, help string) {
	cli.commands = append(cli.commands, Command{fun, name, params, help})
}

func (cli *CLI) PrintCommands() {
	var line string
	var lines []string
	var maximum int

	for _, command := range cli.commands {
		line = command.Name

		for _, param := range command.Params {
			fmt.Print("\x20[" + param + "]")
		}

		if len(line) > maximum {
			maximum = len(line)
		}

		lines = append(lines, line)
	}

	for index, line := range lines {
		fmt.Print("\x20\x20")
		fmt.Print(cli.binary)
		fmt.Print("\x20" + line)
		fmt.Print(strings.Repeat("\x20", maximum-len(line)))
		fmt.Print("\x20" + cli.commands[index].Help)
		fmt.Print("\n")
	}
}

func (cli *CLI) Execute() int {
	flag.Parse()

	query := flag.Arg(0)

	/* default action */
	if query == "" {
		query = "help"
	}

	for _, command := range cli.commands {
		if command.Name == query {
			return command.Function()
		}
	}

	return cli.CallHelp()
}
