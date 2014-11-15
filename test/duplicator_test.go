package scaggold_test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"scaggold"
	"testing"
)

func TestNewDuplicator(t *testing.T) {
	var actual interface{}
	actual = scaggold.NewDuplicator("./a", "./b")

	_, ok := actual.(*scaggold.Duplicator)
	if !ok {
		t.Error("NewDuplicator() must be returns Duplicator struct")
	}
}

func TestDuplicatorRun(t *testing.T) {
	_, filename, _, _ := runtime.Caller(4)
	testDir, _ := filepath.Abs(path.Dir(filename))
	fixture, _ := filepath.Abs(testDir + "/fixtures")
	os.Mkdir(testDir+"/to", 0755)

	// clean up file
	defer func() {
		os.RemoveAll(testDir + "/to")
	}()

	d := scaggold.NewDuplicator(fixture, testDir+"/to")
	// failed
	if !d.Run() {
		t.Error("File duplicate runtime error!")
	}

	// check copied file exists
	if _, err := os.Stat(testDir + "/to/blank.txt"); err != nil {
		t.Error("File duplicate error! fixture:bloank.txt is not exists.")
	}
	if _, err := os.Stat(testDir + "/to/.ignoresample"); err == nil {
		t.Error("File duplicate unexpected! fixture:.ignoresample, dotfile must not be copied.")
	}
}
