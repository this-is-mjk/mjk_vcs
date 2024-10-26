package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/this-is-mjk/mjk/pkg/commands"
	"github.com/this-is-mjk/mjk/pkg/utils"
)

func main() {
	// flags declaration (subcommands)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	commitMessage := commitCmd.String("m", "", "Commit message")
	// input CLI
	args := os.Args[1:]
	// Error if no arguments given
	if len(args) == 0 {
		fmt.Println("No arguments given")
		// yet to implement
		utils.PrintHelp()
		// os.Exit(1)
		return
	}
	switch args[0] {
	case "help":
		fmt.Println("Help")
	case "init":
		commands.Init()
	case "add":
		commands.Add(args[1:])
	case "commit":
		commitCmd.Parse(args[1:])
		commands.Commit(commitMessage)
	case "status":
		commands.Status()
	case "branch":
		commands.Branch(args[1:])
	case "checkout":
		commands.Checkout(args[1])
	case "cat-file":
		commands.CatFile(args[1:])
	case "test":
		commands.Test()
	default:
		fmt.Printf("mjk: not valid command %s\n", args[0])
		utils.PrintHelp()

	}
}
