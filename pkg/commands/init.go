package commands

import (
	"log"

	utils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Init() {
	utils.CreateFile("./.mjk", "HEAD")
	utils.CreateDirectory("./.mjk/objects")	
	log.Printf("\n.mjk initialized\n")
}