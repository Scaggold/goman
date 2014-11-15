package scaggold

import (
	"fmt"
)

type Action struct {
	sys  string
	work string
}

func NewAction(sysDir, workingDir string) *Action {
	return &Action{
		sys:  sysDir,
		work: workingDir,
	}
}

func (a *Action) Execute(args *Arguments) {
	switch cmd, _ := args.GetCommandAt(1); cmd {
	case "get":
		tmpl, ok := args.GetCommandAt(2)
		if !ok {
			fmt.Println("Template is not specified. Plaese type `gm get [template_name]`")
			return
		}

		t := NewTemplate()
		t.Get(tmpl, a.sys+"/"+tmpl)

	case "gen":
		tmpl, ok := args.GetCommandAt(2)
		if !ok {
			fmt.Println("Template is not specified. Plaese type `gm get [template_name]`")
			return
		}

		d := NewDuplicator(a.sys+"/"+tmpl, a.work)
		d.Run()
	}
}
