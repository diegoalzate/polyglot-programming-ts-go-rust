package projector

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Add Operation = iota
	Remove
	Print
)

type ProjectorConfig struct {
	pwd string
	config string
	operation Operation
	args []string
}


func getPwd(opts *ProjectorOpts) (string, error) {
	if opts.Pwd == "" {
		return os.Getwd()
	}

	return opts.Pwd, nil
}

func getConfig(opts *ProjectorOpts) (string, error) {
	if opts.Config == "" {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}

		return path.Join(configDir, "projector", "projector.json"), nil
	}


	return opts.Config, nil
}

func getOperation(opts *ProjectorOpts) Operation {
	if (len(opts.Arguments) == 0) {
		return Print
	}

	if (opts.Arguments[0] == "add") {
		return Add
	}

	if (opts.Arguments[0] == "remove") {
		return Remove
	}

	return Print
}

func getArgs(opts *ProjectorOpts) ([]string, error) {
	operation := getOperation(opts)

	if (operation == Add) {
		if (len(opts.Arguments) != 3) {
			return nil, fmt.Errorf("add expects 2 arguments but received %v", len(opts.Arguments) - 1)
		}

		return opts.Arguments[1:], nil
	}

	if (operation == Remove) {
		if (len(opts.Arguments) != 2) {
			return nil, fmt.Errorf("remove expects 1 argument but received %v", len(opts.Arguments) - 1)
		}

		return opts.Arguments[1:], nil
	}
	

	// assume print
	if (len(opts.Arguments) != 2) {
		return nil, fmt.Errorf("add expects 2 arguments but received %v", len(opts.Arguments) - 1)
	}

	return opts.Arguments[1:], nil
}

func Config(opts *ProjectorOpts) (*ProjectorConfig, error) { 
	pwd, err := getPwd(opts)
	
	if err != nil {
		return nil, err
	}

	config, err := getConfig(opts)

	if err != nil {
		return nil, err
	}

	args, err := getArgs(opts)
	
	if err != nil {
		return nil, err
	}

	return &ProjectorConfig{
		pwd: pwd,
		config: config,
		operation: getOperation(opts),
		args: args,
	}, nil }