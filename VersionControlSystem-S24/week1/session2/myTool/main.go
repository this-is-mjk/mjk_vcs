package main

import (
	"fmt"
	"log"
	"myTool/utils"
	"os"
)

func main() {
	args := os.Args

	// Slicing Arguments string slice because 0th index stores program name
	// args := os.Args[1:]

	// Printing args[0] to show 0th index
	fmt.Println(args[0])

	if len(args) != 3 {
		log.Fatal("Wrong Number of Arguments Arguments")
	}

	switch args[1] {
	case "createDir":
		{
			utils.CreateDir(args[2])
		}
	case "readFile":
		{
			utils.ReadFile(args[2])
		}
  default: 
    {
      fmt.Println("Argument not Available")
      fmt.Println("Available Arguments - createDir, readFile")
    }
		// More cases can be added
	}
}
