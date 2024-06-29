package modles

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	// "github.com/this-is-mjk/mjk/pkg/utils"
	"github.com/this-is-mjk/mjk/pkg/utils"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

type Commit struct {
	Tree          string
	Parent        string
	AuthorName    string
	AuthorEmail   string
	CommiterName  string
	CommiterEmail string
	CommitMessage string
	Time          string
}

func (data *Commit) Decompress(compressedData bytes.Buffer) {
	fileUtils.Decompress(compressedData, data)
}

func GetCommitBody(treeHash string, parent string, commitMessage *string) Commit {

	// set environment in your shell by
	// Example: export AUTHOR_NAME=your_name
	// Execute echo $AUTHOR_NAME to check if the variable has been set in the shell
	// unset AUTHOR_NAME to unset the variable

	commitBody := Commit{
		Tree:          treeHash,
		Parent:        parent,
		AuthorName:    os.Getenv("AUTHOR_NAME"),
		AuthorEmail:   os.Getenv("AUTHOR_EMAIL"),
		CommiterName:  os.Getenv("COMMITTER_NAME"),
		CommiterEmail: os.Getenv("COMMITTER_EMAIL"),
		CommitMessage: *commitMessage,
		Time:          time.Now().String(),
	}
	return commitBody
}

func ShowCommit(commitBody Commit) string {
	// commit body

	// commit (length_of_content)(NULL_BYTE)tree (TREE_ID)
	// parent (PARENT_ID)(NULL_BYTE)
	// author (AUTHOR_NAME) (EMIAL) (TIME_STRING)
	// committer (COMMITTER_NAME) (EMAIL) (TIME_STRING)
	// (COMMIT_MESSAGE)

	returnString := fmt.Sprintf("author %s <%s>\ncommitter %s <%s>\ntime %s\n%s\n",
		commitBody.AuthorName,
		commitBody.AuthorEmail,
		commitBody.CommiterName,
		commitBody.CommiterEmail,
		commitBody.Time,
		commitBody.CommitMessage)
	if commitBody.Parent != "" {
		returnString = fmt.Sprintf("parent %s\n%s", commitBody.Parent, returnString)
	}
	returnString = fmt.Sprintf("tree %s\n%s", commitBody.Tree, returnString)
	return fmt.Sprintf("commit %d\000 %s", len(returnString), returnString)
}

func GetParentCommitId() string {
	// get current branch file, read it and return the parent commit id
	parnentId, _ := fileUtils.ReadFile(utils.GetCurrentBranch())
	if len(parnentId) == 0 {
		return ""
	}
	return string(parnentId)
}

func ReadCommit(commitId string) Commit{
	data, _ := fileUtils.ReadFile(filepath.Join(".", ".mjk", "objects", commitId[:2], commitId[2:]))
	var commit Commit
	commit.Decompress(*bytes.NewBuffer(data))
	return commit
}
