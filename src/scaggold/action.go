package scaggold

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	// get template from repository
	case "get":
		a.RunGet(args)

	// generate template
	case "gen":
		a.RunGenerate(args)

	// show enable templates
	case "list":
		a.RunList(args)

	// purge template
	case "purge":
		a.RunPurge(args)

	// install template from local
	case "install":
		a.RunInstall(args)

	// consiguration
	case "config":
		a.RunConfig(args)

	default:
		fmt.Printf("%s command not found.", cmd)
		ShowHelp()
	}
}

func (a *Action) RunGet(args *Arguments) {
	tmpl, ok := args.GetCommandAt(2)
	if !ok {
		fmt.Println("Template is not specified. Plaese type `gm get [template_name]`")
		return
	}

	t := NewTemplate()
	t.Get(tmpl, a.sys+"/"+tmpl)
}

func (a *Action) RunGenerate(args *Arguments) {
	tmpl, ok := args.GetCommandAt(2)
	if !ok {
		fmt.Println("Template is not specified. Plaese type `gm gen [template_name]`")
		return
	}

	d := NewDuplicator(a.sys+"/"+tmpl, a.work)
	d.Run()
}

func (a *Action) RunList(args *Arguments) {
	list, err := ioutil.ReadDir(a.sys)
	if err != nil {
		fmt.Println("System directory read error.")
		return
	}

	fmt.Println("Installed templates:")
	for _, fi := range list {
		if fi.IsDir() {
			fmt.Printf(" - %s\n", fi.Name())
		}
	}
}

func (a *Action) RunPurge(args *Arguments) {
	tmpl, ok := args.GetCommandAt(2)
	if !ok {
		fmt.Println("Template is not specified. Plaese type `gm purge [template_name]`")
		return
	}

	if _, err := os.Stat(a.sys + "/" + tmpl); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	os.Remove(a.sys + "/" + tmpl)
	fmt.Printf("Template \"%s\" purged.", tmpl)
}

func (a *Action) RunInstall(args *Arguments) {
	tmpl, ok := args.GetCommandAt(2)
	if !ok {
		fmt.Println("Template is not specified. Plaese type `gm install [template_name]`")
		return
	}

	if _, err := os.Stat(a.sys + "/" + tmpl); err == nil {
		fmt.Printf("\"%s\" template is already installed.")
		return
	}

	os.Mkdir(a.sys+"/"+tmpl, 0755)

	d := NewDuplicator(a.work+"/"+tmpl, a.sys+"/"+tmpl)
	d.Silent()

	if !d.Run() {
		fmt.Println("Install failed.")
		return
	}

	fmt.Printf("\"%s\" template installed.\n", tmpl)
}

func (a *Action) RunConfig(args *Arguments) {
	key, ok := args.GetCommandAt(2)
	if !ok {
		fmt.Println("Configuration failed: key must be specified: `gm config [key]`")
		return
	}

	var config map[string]string
	buffer, _ := ioutil.ReadFile(a.sys + "/../config.json")
	json.Unmarshal(buffer, &config)

	if val, ok := args.GetCommandAt(3); ok {
		// setter
		config[key] = val
		if str, err := json.MarshalIndent(config, "  ", "    "); err == nil {
			ioutil.WriteFile(a.sys+"/../config.json", str, 0644)
		}
	} else {
		// getter
		if val, ok := config[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Printf("Undefined config: %s\n", key)
		}
	}
}
