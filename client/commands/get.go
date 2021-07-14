package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/tenntenn/natureremo"
)

type getCommand struct {
}

func NewGetCommand() cli.Command {
	return &getCommand{}
}

func (c *getCommand) Help() string {
	return "usage: remo get <hostname>"
}

func (c *getCommand) Synopsis() string {
	return "Get the signal that was sent to the device."
}

func (c *getCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stdout, c.Help())
		return 1
	}

	hostname := args[0]
	client := natureremo.NewLocalClient(hostname)
	result, err := client.Fetch(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 2
	}

	err = json.NewEncoder(os.Stdout).Encode(*result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 2
	}

	return 0
}
