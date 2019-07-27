package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cixtor/slackapi"
)

// Function defines a CLI method.
type Function func() int

// CLI defines the core of the program.
type CLI struct {
	binary   string
	api      *slackapi.SlackAPI
	commands []Command
}

// Command defines an option to call an API method.
type Command struct {
	Function Function
	Name     string
	Params   []string
	Help     string
}

// NewCLI returns an instance of CLI.
func NewCLI(binary string) *CLI {
	cli := new(CLI)

	cli.binary = binary
	cli.api = slackapi.New()

	return cli
}

// AutoAuthenticate searches for a token among the environment variablles to
// authenticate the HTTP requests with the web API service.
func (cli *CLI) AutoAuthenticate() {
	cli.api.SetToken(os.Getenv("SLACK_TOKEN"))
	cli.api.SetCookie(os.Getenv("SLACK_COOKIE"))
}

// Register adds support for a new command.
func (cli *CLI) Register(fun Function, name string, params []string, help string) {
	cli.commands = append(cli.commands, Command{fun, name, params, help})
}

// Execute calls a method to send a HTTP request to the web API service.
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

// PrintCommands builds the usage options for the help command.
func (cli *CLI) PrintCommands() {
	var line string
	var lines []string
	var maximum int

	for _, command := range cli.commands {
		line = command.Name

		for _, param := range command.Params {
			line += fmt.Sprint("\x20[" + param + "]")
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

// PrintJSON takes a generic struct with the response from the web API service
// and then proceeds to encode it as a JSON string. It uses the JSON string to
// obtain the status and possible errors from the HTTP request by re-decoding
// into another smaller strust. It returns an Unix exit code representing the
// success or failure of the operation, zero and one respectively.
func (cli *CLI) PrintJSON(v interface{}) int {
	var res slackapi.Response

	out, err := json.MarshalIndent(v, "", "\x20\x20")

	if err != nil {
		fmt.Printf("{\"ok\":false, \"error\":\"json.encode; %s\"}\n", err.Error())
		return 1
	}

	if err := json.NewDecoder(bytes.NewReader(out)).Decode(&res); err != nil {
		fmt.Printf("{\"ok\":false, \"error\":\"json.decode; %s\"}\n", err.Error())
		return 1
	}

	if !res.Ok && res.Error != "" {
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			log.Fatalln("json.encode;", err)
		}

		return 1
	}

	fmt.Printf("%s\n", out)
	return 0
}

// Number attempts to decode the user input as an integer.
func (cli *CLI) Number(index int, initial int) int {
	input := flag.Arg(index)

	if input == "" {
		return initial
	}

	if number, err := strconv.Atoi(input); err == nil {
		return number
	}

	return initial
}
