package utils

import (
	"fmt"
  "log"
	"os"
)

func ReadFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Println(string(file))  
}
