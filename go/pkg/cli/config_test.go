package projector_test

import (
	"testing"

	projector "github.com/diegoalzate/polyglot-programming-ts-go-rust/pkg/cli"
)


func TestConfigPrint(t *testing.T) {
	opts := &projector.ProjectorOpts{
		Pwd: "",
		Config: "",
		Arguments: []string{"print", "b"},
	}

	config, err := projector.Config(opts)

	if err != nil {
		t.Fatalf("FAILED TO CREATE CONFIG %v", err)
	}

	if config.Operation != projector.Print {
		t.Fatalf("FAILED TO CREATE OPERATION PRINT %v",  config.Operation)
	}
}

func TestConfigAdd(t *testing.T) {
	opts := &projector.ProjectorOpts{
		Pwd: "",
		Config: "",
		Arguments: []string{"add", "key", "valie"},
	}

	config, err := projector.Config(opts)

	if err != nil {
		t.Fatalf("FAILED TO CREATE CONFIG %v", err)
	}

	if config.Operation != projector.Add {
		t.Fatalf("FAILED TO CREATE OPERATION ADD %v",  config.Operation)
	}
}

func TestConfigRemove(t *testing.T) {
	opts := &projector.ProjectorOpts{
		Pwd: "",
		Config: "",
		Arguments: []string{"remove", "key"},
	}

	config, err := projector.Config(opts)

	if err != nil {
		t.Fatalf("FAILED TO CREATE CONFIG %v", err)
	}

	if config.Operation != projector.Remove {
		t.Fatalf("FAILED TO CREATE OPERATION REMOVE %v",  config.Operation)
	}
}