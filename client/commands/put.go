package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		os.Stdout.WriteString(fmt.Sprintln(c.Help()))
		return 1
	}

	hostname := args[0]
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		return 2
	}

	var signal *natureremo.IRSignal
	err = json.Unmarshal(data, &signal)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		return 2
	}

	client := natureremo.NewLocalClient(hostname)
	err = client.Emit(context.Background(), signal)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		return 2
	}

	os.Stderr.WriteString("Successfully sent!\n")
	return 0
}
