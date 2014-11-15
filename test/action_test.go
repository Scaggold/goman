package scaggold_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"scaggold"
	"testing"
)

func TestNewAction(t *testing.T) {
	var actual interface{}
	actual = scaggold.NewAction("foo", "bar")

	_, ok := actual.(*scaggold.Action)
	if !ok {
		t.Error("NewAction() must be returns Action struct")
	}
}

func TestExecute(t *testing.T) {
	_, filename, _, _ := runtime.Caller(4)
	testDir, _ := filepath.Abs(path.Dir(filename))
	os.MkdirAll(testDir+"/tmp/template", 0755)
	os.MkdirAll(testDir+"/tmp/target", 0755)

	fmt.Println(testDir)

	// clean up file
	defer func() {
		os.RemoveAll(testDir + "/tmp")
	}()

	fakeGetArgs := []string{"get", "sample"}
	fakeGenArgs := []string{"gen", "sample"}

	a := scaggold.NewAction(testDir+"/tmp/template", testDir+"/tmp/target")
	args := scaggold.NewArguments()
	args.Parse(fakeGetArgs)

	// get action
	a.Execute(args)
	stat, err := os.Stat(testDir + "/tmp/template/sample")
	if err != nil || !stat.IsDir() {
		t.Error("get action failed!")
	}

	args = scaggold.NewArguments()
	args.Parse(fakeGenArgs)

	a.Execute(args)
	if _, err := os.Stat(testDir + "/tmp/target/index.js"); err != nil {
		t.Error("gen action failed!")
	}
}
