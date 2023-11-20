package projector_test

import (
	"fmt"
	"testing"

	projector "github.com/diegoalzate/polyglot-programming-ts-go-rust/pkg/cli"
)

func getConfig(pwd string) projector.ProjectorConfig {
	return projector.ProjectorConfig{
		Pwd: pwd,
		Config: "/foo/bar/baz",
		Operation: projector.Add,
		Args: []string{},
	}
}

func getData() projector.ProjectorData {
	return projector.ProjectorData{
		Projector: map[string]map[string]string{
			"/foo/bar/baz/buzz": {
				"foo": "bar1",
			  },
			  "/foo/bar/baz": {
				"foo": "bar2",
			  },
			  "/foo/bar": {
				"foo": "bar3",
			  },
			  "/foo": {
				"foo": "bar4",
			  },
			  "/": {
				"foo": "bar5",
				"bar": "bazz1",
			  },
		},
	}
}

func createDefaultProjector(pwd string, data projector.ProjectorData) *projector.Projector {
	return &projector.Projector{
		Config: getConfig(pwd),
		Data: data,
	}
}

func TestGetValue(t *testing.T) {
	projector := createDefaultProjector("/foo/bar", getData())
	value, found := projector.GetValue("foo")

	if !found {
		t.Error("Did not receive value on GetAll")
	}

	if value != "bar3" {
		t.Errorf("expected value %v but received %v", "bar3", value)
	}

	value, found = projector.GetValue("bar")

	if !found {
		t.Error("Did not receive value on GetAll")
	}

	if value != "bazz1" {
		t.Errorf("expected value %v but received %v", "bazz1", value)
	}
}

func TestGetAllValues(t *testing.T) {
	projector := createDefaultProjector("/foo/bar/baz", getData())
	value := projector.GetValueAll()
	fmt.Printf("%v", value)
	if value["foo"] != "bar2" {
		t.Errorf("foo key expected value %v but received %v", "bar2", value)
	}

	if value["bar"] != "bazz1" {
		t.Errorf("bar key expected value %v but received %v", "bazz1", value)
	}
}

func TestSetValue(t *testing.T) {
	projector := createDefaultProjector("/foo/bar/baz", getData())
	value, found := projector.GetValue("foo")

	if !found {
		t.Error("Did not receive value on GetValue")
	}

	if value != "bar2" {
		t.Errorf("expected value %v but received %v", "bar2", value)
	}

	projector.SetValue("foo", "betterThanBar3")

	value, found = projector.GetValue("foo")

	if !found {
		t.Error("Did not receive value on GetValue")
	}

	if value != "betterThanBar3" {
		t.Errorf("expected value %v but received %v", "betterThanBar3", value)
	}

	projector = createDefaultProjector("/foo", getData())	
	
	value, found = projector.GetValue("foo")

	if !found {
		t.Error("Did not receive value on GetValue")
	}

	if value != "bar4" {
		t.Errorf("expected value %v but received %v", "bar4", value)
	}
}

func TestRemoveValue(t *testing.T) {
	projector := createDefaultProjector("/foo/bar/baz", getData())
	value, found := projector.GetValue("foo")

	if !found {
		t.Error("Did not receive value on GetValue")
	}

	if value != "bar2" {
		t.Errorf("expected value %v but received %v", "bar2", value)
	}

	projector.RemoveValue("foo")

	value, found = projector.GetValue("foo")

	if value == "bar2" {
		t.Errorf("expected value to not exist but received %v", value)
	}

	if value != "bar3" {
		t.Errorf("expected value %v but received %v", "bar3", value)
	}
}