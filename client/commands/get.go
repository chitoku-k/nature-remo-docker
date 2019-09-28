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
		os.Stdout.WriteString(fmt.Sprintln(c.Help()))
		return 1
	}

	hostname := args[0]
	client := natureremo.NewLocalClient(hostname)
	result, err := client.Fetch(context.Background())
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		return 2
	}

	value, err := json.Marshal(*result)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		return 2
	}

	os.Stdout.WriteString(fmt.Sprintf("%s\n", value))
	return 0
}
