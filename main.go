package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/agusjanardana/routehash/hashRoute"
)

func main() {
	// define flags
	nameForRoute := flag.String("name", "", "route name hashed with sha256")
	method := flag.String("method", "", "route method")
	apiGroup := flag.String("apiGroup", "", "apiGroup")
	controller := flag.String("controller", "", "add your controller")

	// parse flags
	flag.Parse()

	// check for custom command
	if *nameForRoute == "" || *method == "" || *apiGroup == "" || *controller == "" {
		fmt.Println("Error: no command provided, please fill all command needed")
		dataNeeded := "routehash --name=\"your_route_name\" method=\"method_name\" apiGroup=\"your_api_group\" controller=\"your_controller\""
		fmt.Println(dataNeeded)
		os.Exit(1)
	}

	// hash
	dataToReturn := hashRoute.HashRoute(*nameForRoute)

	// write it to app/configs/routes
	hashRoute.WriteToRoute2(*apiGroup, dataToReturn, *method, *controller)

	// run custom command
	fmt.Println("Route name hashed for", *nameForRoute, "is :", dataToReturn)
	fmt.Println("Route has been written in app/configs/route.go")
}
