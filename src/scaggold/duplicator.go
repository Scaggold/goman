package scaggold

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Duplicator struct {
	fromPath  string
	toPath    string
	fromRegex *regexp.Regexp
	verbose   bool
	runnables []string
}

// Utility function
// Traling slash on right segument
func TrailSlash(path string) string {
	return strings.TrimRight(path, "/") + "/"
}

func NewDuplicator(fromPath, toPath string) *Duplicator {
	return &Duplicator{
		fromPath:  TrailSlash(fromPath),
		toPath:    TrailSlash(toPath),
		fromRegex: regexp.MustCompile("^" + TrailSlash(fromPath)),
		verbose:   true,
	}
}

func (d *Duplicator) Silent() {
	d.verbose = false
}

func (d *Duplicator) Run() bool {
	err := filepath.Walk(d.fromPath, d.walkFunc)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	for _, file := range d.runnables {
		fp, _ := os.Open(file)
		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		shebang := []rune(scanner.Text())
		if string(shebang[0]) == "#" && string(shebang[1]) == "!" {
			command := string(shebang[2:])
			cmd := NewCommand(command)
			cmd.Arg(file)
			cmd.Exec()
		} //cmd := NewCommand("bash")
		//cmd.SendCommand("cd " + d.toPath)
		//cmd.SendCommand(file)
		//cmd.Exec()
	}

	return true
}

func (d *Duplicator) walkFunc(path string, info os.FileInfo, err error) error {
	rel := d.fromRegex.ReplaceAll([]byte(path), []byte(""))
	dest := d.toPath + string(rel)

	// Skip dotfile
	if len(rel) > 0 && string(rel[0]) == "." {
		if string(rel) == ".install" {
			d.runnables = append(d.runnables, path)
		}
		return nil
	}

	// Does walking directory?
	if info.IsDir() {
		os.Mkdir(dest, info.Mode())
		return nil
	}

	// Copy file
	if buffer, err := ioutil.ReadFile(path); err == nil {
		if err := ioutil.WriteFile(dest, buffer, info.Mode()); err == nil {
			if d.verbose {
				fmt.Printf("%s ====> %s\n", path, dest)
			}
			return nil
		} else {
			fmt.Printf("%v\n", err)
			return errors.New("File write error")
		}
	} else {
		fmt.Printf("%v\n", err)
		return errors.New("File read error")
	}

	return nil
}
