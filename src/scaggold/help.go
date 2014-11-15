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
get [template_name]      - Install template from scaggold repository
gen [template_name]      - Generate template to current directory
purge [template_name]    - Remove installed template
install [directort_path] - Install local template
config [keyname]         - Get configuration value
config [keyname] [value] - Set configuration value
list                     - Show installed template list
help                     - Show this help

Options:
-h    - Show this help`

	fmt.Println(help)
}
