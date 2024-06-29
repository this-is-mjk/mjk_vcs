package fileUtils

import (
	"os"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// fmt.Println("Error in reading the file")
		return nil, err
	}
	return data, nil
}
