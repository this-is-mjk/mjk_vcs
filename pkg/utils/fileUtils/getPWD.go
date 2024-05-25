package fileUtils

import (
	"log"
	"os"
)

func Get_pwd() string {
	// returns working dir in string
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to fetch present working directory\n ERROR: %s", err)
	}
	return wd
}
