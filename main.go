package main

import (
	"flag"
	"fmt"
	"os"
	"routehash/hashRoute"
)

func main() {
	// define flags
	customCommand := flag.String("name", "", "route name hashed with sha256")

	// parse flags
	flag.Parse()

	// check for custom command
	if *customCommand == "" {
		fmt.Println("Error: no command provided")
		os.Exit(1)
	}

	// hash
	dataToReturn := hashRoute.HashRoute(*customCommand)

	// run custom command
	fmt.Println("Route name hashed for",*customCommand,"is :", dataToReturn)
}