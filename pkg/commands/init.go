package commands

import (
	"bytes"
	"log"

	utils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Init() {
	// create .mjk directory and HEAD file
	utils.CreateFile("./.mjk", "HEAD")
	// write the pointer to main pointer file in HEAD
	utils.WriteFile("./.mjk/HEAD", *bytes.NewBufferString("ref: refs/heads/main"))
	// create objects directory
	utils.CreateDirectory("./.mjk/objects")
	// create refs directory, heads directory and inside that main file
	utils.CreateFile("./.mjk/refs/heads", "main")
	log.Printf("\n.mjk initialized\n")
}
