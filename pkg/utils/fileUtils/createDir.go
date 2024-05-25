package fileUtils

import (
	"fmt"
	"os"
)

func CreateDirectory(path string) {
	// Create a directory using path and name
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Printf("Unable to create %s\nERROR:- %s", path, err)
	}
}