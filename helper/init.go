package helper

import (
	"log"
	"os"
)

func CreateDir(name string) {
	os.MkdirAll(name, os.ModePerm)
	log.Printf("New Directory got created with name %s", name)
}

func Init() {
	// // returns working dir in string
	// wd, err := os.Getwd()
	// if err != nil {
	// 	return nil, err
	// }
	CreateDir(".mjk/HEAD")
	CreateDir(".mjk/objects")
}
