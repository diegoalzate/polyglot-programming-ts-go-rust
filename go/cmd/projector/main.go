package main

import (
	"fmt"
	"log"

	opts "github.com/diegoalzate/polyglot-programming-ts-go-rust/pkg/cli"
)

func main() {
	opts, err := opts.GetOptions()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("opts: %+v", opts)
}