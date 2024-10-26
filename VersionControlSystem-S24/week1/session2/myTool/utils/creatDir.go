package utils

import (
	"log"
	"os"
)

func CreateDir(name string) {
  os.Mkdir(name, os.ModePerm)
  log.Printf("New Directory got created with name %s", name)
}
