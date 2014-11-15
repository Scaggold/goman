package scaggold

import (
	"fmt"
	"os"
	"os/exec"
)

const REPOSITORY string = "git@github.com:Scaggold/"

type Template struct{}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) Get(name, templateDir string) bool {
	commandArgs := []string{
		"clone",
		REPOSITORY + name + ".git",
		templateDir,
	}
	cmd := exec.Command("git", commandArgs...)
	cmd.Stdout = os.Stdout

	fmt.Printf("Getting template \"%s\" from %s..", name, REPOSITORY+name+".git")

	if err := cmd.Run(); err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	fmt.Println("done!")
	return true
}
