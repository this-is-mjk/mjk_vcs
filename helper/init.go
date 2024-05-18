package helper

import (
	"log"
	"os"
)
func Init() {
	CreateDir(".mjk");
	os.Create(".mjk/HEAD");
	CreateDir(".mjk/objects");
	log.Printf("\n.mjk initialized\n");
}
