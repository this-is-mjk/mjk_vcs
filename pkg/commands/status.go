package commands

import (
	"fmt"
	"path/filepath"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Status() {
	// get the current branch
	branch := utils.GetCurrentBranch()
	fmt.Printf("On branch %s\n", branch[16:])
	// TODO: find the not added and not commited files.

	// find changes to be committed
	// read the commit, get the tree id
	lastCommitId, _ := fileUtils.ReadFile(filepath.Join(".", branch))
	treeId := modles.ReadCommit(string(lastCommitId)).Tree
	// read tree
	modles.ShowTree(modles.ReadTree(treeId))

	// read the commit id form the current pointer
	// read the commit
	// decompress the commit body, make a commit struct now, change it also in commit package
	// make blob struct, tree struct
	// itrate over all the files in the directory, check it exist in index file, if exist then check the hash of the file with the hash in the index file, then check the hash of the file with the hash in the commit file,
}
