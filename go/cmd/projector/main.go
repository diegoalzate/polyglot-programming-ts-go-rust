package main

import (
	"fmt"
	"log"

	cli "github.com/diegoalzate/polyglot-programming-ts-go-rust/pkg/cli"
)

func main() {
	opts, err := cli.GetOptions()

	if err != nil {
		log.Fatal(err)
	}

	config, err := cli.Config(opts);


	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("opts: %+v", opts)
	fmt.Printf("config %+v", config)
}