package main

import (
	"fmt"
	"os"
)

func main() {
	// cutting slice from 1 to end
	// because Args[0] stores binary name itself
	args := os.Args[1:]

	// If no args are provided
	if len(args) == 0 {
		// you can also display help section when no args provided
		fmt.Printf("your_vcs_name: no args provided\n")
		// exit binary execution with exit code 1
    // 0 is considered success code by convention
		os.Exit(1)
	}

	command := args[0]
	switch command {
	case "init":
		//InitHandler()
	case "commit":
		//CommitHandler()
	default:
		fmt.Printf("your_vcs_name: %s is not a valid command\n", command)
		os.Exit(1)
	}

}
