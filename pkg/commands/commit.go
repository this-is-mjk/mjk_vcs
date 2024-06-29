package commands

import (
	"bytes"
	"fmt"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils"
	creator "github.com/this-is-mjk/mjk/pkg/utils/componentCreater"
	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Commit(commitMessage *string) {
	PWD := fileUtils.Get_pwd()

	tree, isEmpty := creator.Tree(PWD, creator.ReadIndex())
	if isEmpty {
		fmt.Println("Nothing to commit, please add files to commit")
		return
	}
	commit := modles.GetCommitBody(tree, modles.GetParentCommitId(), commitMessage)
	modles.ShowCommit(commit)
	commit_sha1 := utils.Sha1(commit)
	creator.Object(commit_sha1, fileUtils.Compress(commit))
	// write the commit id to head
	var buffer bytes.Buffer
	buffer.WriteString(commit_sha1)
	// get the current branch and write the commit id to the branch file
	fileUtils.WriteFile(utils.GetCurrentBranch(), buffer)
	modles.ShowCommit(commit)
}
