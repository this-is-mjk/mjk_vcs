package commands

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Checkout(branch string) {
	// check if the branch exists
	readFile, err := fileUtils.ReadFile(filepath.Join(".", ".mjk", "refs", "heads", branch))
	if os.IsNotExist(err) {
		fmt.Printf("%v branch does not exist\n", branch)
		os.Exit(1)
	}
	branchLastCommitPointer := filepath.Join(".mjk", string(readFile))[5:]
	fmt.Print(branchLastCommitPointer)
	fileUtils.WriteFile("./.mjk/HEAD", *bytes.NewBufferString("ref: refs/heads/" + branch))
}
