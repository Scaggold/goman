package scaggold

import (
	"bufio"
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

	reader := bufio.NewReader(cmdOut)

	fmt.Println("Running install sctipt...")

	for _, c := range c.shells {
		io.WriteString(writer, c+"\n")
	}

	var buf []byte
	for {
		if num, err := reader.Read(buf); err == io.EOF || num == 0 {
			cmdOut.Close()
			break
		}
		fmt.Println(buf)
	}
}
