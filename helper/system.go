package helper

import (
	"log"
	"os"
	"path/filepath"
)

func Is_first_commit() bool {
	f, err := os.Stat(filepath.Join(Get_pwd(), ".mjk", "HEAD"))
	if err != nil {
		log.Fatalf("Unable to fetch HEAD\nERROR:- %s", err)
	}
	return f.Size() == 0
}
func CreateDir(name string) {
	if err := os.MkdirAll(name, os.ModePerm); err != nil {
		log.Fatalf("Unable to create directory\nERROR:- %s", err)
	}
}
func Get_pwd() string {
	// returns working dir in string
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to fetch present working directory\n ERROR: %s", err)
	}
	return wd
}
