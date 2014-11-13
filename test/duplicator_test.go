package scaggold_test

import (
	"github.com/Scaggold/goman"
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

// TODO: need filesystem  test
