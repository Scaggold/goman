package scaggold

import (
	"fmt"
	"io"
	"os/exec"
)

type Command struct {
	command string
	shells  []string
}

func NewCommand(cmd string) *Command {
	return &Command{
		command: cmd,
	}
}

func (c *Command) SendCommand(command string) {
	c.shells = append(c.shells, command)
}

func (c *Command) Exec() {
	cmd := exec.Command(c.command)

	writer, _ := cmd.StdinPipe()
	cmdOut, _ := cmd.StdoutPipe()

	cmd.Start()

	fmt.Println("Running install sctipt...")

	for _, c := range c.shells {
		io.WriteString(writer, c+"\n")
	}

	cmdOut.Close()
}
