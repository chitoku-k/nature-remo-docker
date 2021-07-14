package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/tenntenn/natureremo"
)

type putCommand struct {
}

func NewPutCommand() cli.Command {
	return &putCommand{}
}

func (c *putCommand) Help() string {
	return "usage: remo put <hostname>"
}

func (c *putCommand) Synopsis() string {
	return "Put the given signal to the device."
}

func (c *putCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stdout, c.Help())
		return 1
	}

	hostname := args[0]

	var signal *natureremo.IRSignal
	err := json.NewDecoder(os.Stdin).Decode(&signal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 2
	}

	client := natureremo.NewLocalClient(hostname)
	err = client.Emit(context.Background(), signal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 2
	}

	fmt.Fprintln(os.Stderr, "Successfully sent!")
	return 0
}
