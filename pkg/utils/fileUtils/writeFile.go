package fileUtils

import (
	"bytes"
	"fmt"
	"os"
)

func WriteFile(path string, data bytes.Buffer) {
	// Write data to a file
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Unable to create %s\nERROR:- %s", path, err)
	}
	defer file.Close()
	if _, err := data.WriteTo(file); err != nil {
		fmt.Printf("Error writing data to file\nERROR:- %s", err)
	}
}
