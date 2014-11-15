package scaggold_test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"scaggold"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	var actual interface{}
	actual = scaggold.NewTemplate()

	_, ok := actual.(*scaggold.Template)
	if !ok {
		t.Error("NewTemplate() must be returns Template struct")
	}
}

func TestRun(t *testing.T) {
	_, filename, _, _ := runtime.Caller(4)
	testDir, _ := filepath.Abs(path.Dir(filename))
	os.Mkdir(testDir+"/tmp", 0755)

	// clean up file
	defer func() {
		os.RemoveAll(testDir + "/tmp")
	}()

	tmpl := scaggold.NewTemplate()
	if !tmpl.Get("sample", testDir+"/tmp/sample") {
		t.Error("Template get runtime error!")
	}

	// check copied file exists
	stat, err := os.Stat(testDir + "/tmp/sample")
	if err != nil {
		t.Error("Template get error: directory not exists.")
	}
	if !stat.IsDir() {
		t.Error("Template get error: template is not a directory.")
	}

}
