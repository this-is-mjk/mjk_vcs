package fileUtils

import (
	"fmt"
	"os"
)

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file\nERROR:- %s", err)
	}
	return data
}
