package main

import (
	"encoding/json"
	"fmt"
	"log"

	cli "github.com/diegoalzate/polyglot-programming-ts-go-rust/pkg/cli"
)

func main() {
	opts, err := cli.GetOptions()

	if err != nil {
		log.Fatal(err)
	}

	config, err := cli.Config(opts)

	if err != nil {
		log.Fatal(err)
	}

	proj, err := cli.FromConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	if proj.Config.Operation == cli.PrintAll {
		values := proj.GetValueAll()
		bytesJson, err := json.Marshal(values)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v", string(bytesJson))
	}

	if proj.Config.Operation == cli.Print {
		value, found := proj.GetValue(proj.Config.Args[0])

		if found {
			fmt.Printf("%v", value)
		}
	}

	if proj.Config.Operation == cli.Add {
		proj.SetValue(proj.Config.Args[0], proj.Config.Args[1])
		proj.Save()
	}

	if proj.Config.Operation == cli.Remove {
		proj.RemoveValue(proj.Config.Args[0])
		proj.Save()
	}
}
