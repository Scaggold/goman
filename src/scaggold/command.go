package scaggold

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	command string
	args    []string
}

func NewCommand(cmd string) *Command {
	split := strings.Split(cmd, " ")
	command := &Command{
		command: split[0],
	}
	command.Arg(split[1:]...)

	return command
}

func (c *Command) Arg(arg ...string) {
	for _, a := range arg {
		c.args = append(c.args, a)
	}
}

func (c *Command) Exec() {
	cmd := exec.Command(c.command, c.args...)
	stdout := bytes.NewBuffer(nil)

	cmd.Stdout = stdout

	defer func() {
		out := stdout.Bytes()
		fmt.Print(string(out))
	}()

	fmt.Println("Running install sctipt...")

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
