package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"scaggold"
)

func main() {
	// parse arguments
	args := scaggold.NewArguments()
	args.Alias("h", "help", nil)
	args.Parse(os.Args[1:])

	home := os.Getenv("HOME")
	sysDir := home + "/.goman"
	if home == "" {
		fmt.Println("$HOME directory is not found. Plaese check your OS environment...")
		return
	}

	// help check
	if isNeedShowHelp(args) {
		scaggold.ShowHelp()
		return
	}

	// check initialize
	if _, err := os.Stat(sysDir); err != nil {
		fmt.Print("Initializing goman tools...")
		if !initialize(home) {
			fmt.Println("Failed!")
			return
		} else {
			fmt.Println("done!")
		}
	}

	// execute
	cwd, _ := os.Getwd()
	action := scaggold.NewAction(sysDir+"/templates", cwd)
	action.Execute(args)
}

func isNeedShowHelp(args *scaggold.Arguments) bool {
	if args.GetCommandSize() == 0 {
		return true
	}

	if _, ok := args.GetOption("help"); ok {
		return true
	}

	if help, _ := args.GetCommandAt(1); help == "help" {
		return true
	}

	return false
}

func initialize(home string) bool {
	path := home + "/.goman"
	if err := os.Mkdir(path, 0755); err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	// setting file create
	if err := ioutil.WriteFile(path+"/config.json", []byte("{}"), 0644); err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	// template directory create
	if err := os.Mkdir(path+"/templates", 0755); err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	return true
}
