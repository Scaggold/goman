package scaggold

import (
	"fmt"
)

func ShowHelp() {
	help := `goman, A template scaffold generator
====================================
Usage:
$ gm [subcommand,...] [options,...]

Subcommands:
get   - Install template from scaggold repository
gen   - Generate template to current directory
purge - Remove installed template
help  - Show this help

Options:
-h    - Show this help`

	fmt.Println(help)
}
