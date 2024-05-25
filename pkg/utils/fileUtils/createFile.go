package fileUtils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path string, name string) {
	CreateDirectory(path)
	// Path till the directory where file is to be created
	file, err := os.Create(filepath.Join(path, name)) //create a new file
	if err != nil {
		fmt.Printf("Unable to create %s\nERROR:- %s", path, err)
		return
	}
	defer file.Close()
}
