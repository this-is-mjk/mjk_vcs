package commands

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils"
	creator "github.com/this-is-mjk/mjk/pkg/utils/componentCreater"
	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func get_parent() string {
	PWD := fileUtils.Get_pwd()
	parent_id := fileUtils.ReadFile(filepath.Join(PWD, ".mjk", "HEAD"))
	// fmt.Println(len(parent_id))
	if len(parent_id) != 0 {
		return fmt.Sprintf("parent %s\n", parent_id)
	}
	return ""
}
func Commit(commitMessage *string) {
	PWD := fileUtils.Get_pwd()

	// commit body

	// commit (length_of_content)(NULL_BYTE)tree (TREE_ID)
	// parent (PARENT_ID)(NULL_BYTE)
	// author (AUTHOR_NAME) (EMIAL) (TIME_STRING)
	// committer (COMMITTER_NAME) (EMAIL) (TIME_STRING)
	// (COMMIT_MESSAGE)

	// set environment in your shell by
	// Example: export AUTHOR_NAME=your_name
	// Execute echo $AUTHOR_NAME to check if the variable has been set in the shell
	// unset AUTHOR_NAME to unset the variable

	commit := fmt.Sprintf("tree %s\n%sauthor %s %s\ncommitter %s %s\n%s\n",
		creator.Tree(PWD),
		get_parent(),
		os.Getenv("AUTHOR_NAME"),
		os.Getenv("AUTHOR_EMAIL"),
		os.Getenv("COMMITTER_NAME"),
		os.Getenv("COMMITTER_EMAIL"),
		*commitMessage)

	fmt.Println(commit)
	commit_sha1 := hex.EncodeToString(utils.Sha1(commit))
	creator.Object(commit_sha1, fileUtils.Compress(commit))
	// write the commit id to head
	var buffer bytes.Buffer
	buffer.WriteString(commit_sha1)
	fileUtils.WriteFile(filepath.Join(PWD, ".mjk", "HEAD"), buffer)
}
