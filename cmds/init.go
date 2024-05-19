package cmds

import (
	"log"
	"mjk_vcs/helper"
	"os"
)

func Init() {
	helper.CreateDir(".mjk")
	os.Create(".mjk/HEAD")
	helper.CreateDir(".mjk/objects")
	log.Printf("\n.mjk initialized\n")
}
