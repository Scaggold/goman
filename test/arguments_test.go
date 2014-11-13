package scaggold_test

import (
	"github.com/Scaggold/goman"
	"testing"
)

func TestNewArguments(t *testing.T) {
	var actual interface{}
	actual = scaggold.NewArguments()

	_, ok := actual.(*scaggold.Arguments)
	if !ok {
		t.Error("NewArguments() must be returns Arguments struct")
	}
}

func TestAliasWithNextValue(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f", "255"})

	if v, ok := args.GetOption("foo"); ok {
		if v != "255" {
			t.Errorf("Aliased flag parse error. Input is 255, actucal %s", v)
		}
	} else {
		t.Errorf("Aliase set, but not parsing")
	}
}

func TestAliasWithNoValue(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	if v, ok := args.GetOption("foo"); ok {
		if v != true {
			t.Error("Aliased flag parse error. f flag is not set")
		}
	} else {
		t.Errorf("Aliase set, but not parsing")
	}
}

func TestAliasWithValue(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-fbar"})

	if v, ok := args.GetOption("foo"); ok {
		if v != "bar" {
			t.Error("Aliased flag parse error. Input is bar, actual %s", v)
		}
	} else {
		t.Errorf("Aliase set, but not parsing")
	}
}

func TestGetOptionExists(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	if v, ok := args.GetOption("foo"); ok {
		if v != true {
			t.Errorf("GetOption error. Input is none, actual %s", v)
		}
	} else {
		t.Errorf("Aliase set, but not parsing")
	}
}

func TestGetOptionNotExists(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	if _, ok := args.GetOption("baz"); ok {
		t.Error("Unknown option supplied, but exists")
	}
}

func TestGetCommandExists(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	if cmd, ok := args.GetCommandAt(1); !ok || cmd != "foo" {
		t.Errorf("Command test error, assert foo, actual %s", cmd)
	}
}

func TestGetCommandNotExists(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{})

	if _, ok := args.GetCommandAt(1); ok {
		t.Error("Command not suppiled, but exists")
	}
}

func TestGetCommands(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	cmds := args.GetCommands()
	if cmds[0] != "foo" || cmds[1] != "bar" {
		t.Errorf("Command test error, command list asset error")
	}
}
func TestGetCommandSize(t *testing.T) {
	args := scaggold.NewArguments()

	args.Alias("f", "foo")
	args.Parse([]string{"foo", "bar", "-f"})

	if size := args.GetCommandSize(); size != 2 {
		t.Errorf("Command test error, command size not match")
	}
}
