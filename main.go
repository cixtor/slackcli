package main

const binary = "slackcli"
const version = "0.0.1"

func main() {
	cli := NewCLI(binary)

	cli.Register(cli.CallApiTest, "api.test", []string{}, "Checks API calling code")
	cli.Register(cli.CallAuthTest, "auth.test", []string{}, "Checks authentication and identity")
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	cli.Execute()
}
