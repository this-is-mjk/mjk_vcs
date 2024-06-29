package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Branch(branches []string) {
	for _, branch := range branches {
		fileUtils.CreateFile("./.mjk/refs/heads", branch)
		mainPointer, _ := fileUtils.ReadFile("./.mjk/refs/heads/main")
		if string(mainPointer) == "" {
			fmt.Print("no commits on main branch yet to form barnch")
			os.Exit(1)
		}
		fileUtils.WriteFile("./.mjk/refs/heads/"+branch, *bytes.NewBuffer(mainPointer))
	}
}
