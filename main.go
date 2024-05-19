package main

import (
	"fmt"
	"mjk_vcs/cmds"
	"mjk_vcs/helper"
	"os"
)

func main() {
	// input cla
	args := os.Args[1:]
	// Error if no arguments given
	if len(args) != 2 {
		fmt.Println("Invalid Use")
		// yet to implement
		helper.PrintHelp()
		os.Exit(1)
	}
	if args[0] == "mjk" {
		switch args[1] {
		case "help":
			helper.PrintHelp()
		case "init":
			cmds.Init()
		case "commit":
			cmds.Commit()
		default:
			fmt.Printf("mjk: not valid command %s\n", args[1])
			helper.PrintHelp()
		}
	}

}
