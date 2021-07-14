package main

import (
	"fmt"
	"os"

	"github.com/chitoku-k/natureremo-cli/commands"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("natureremo-client", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return commands.NewGetCommand(), nil
		},
		"put": func() (cli.Command, error) {
			return commands.NewPutCommand(), nil
		},
	}

	status, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	os.Exit(status)
}
