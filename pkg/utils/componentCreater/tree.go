package creator

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Tree(dirPath string, indexData modles.IndexFile) (string, bool) {
	isEmpty := true
	var tree string = ""
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		// this function inside filepath.WaldDir is called for each file and directory
		// the path is the path of the file or directory and it will start form PWD/.
		// as called form commit.go

		// only walk on top level
		parent := filepath.Dir(path)
		if parent != dirPath {
			return nil
		}
		// return if default files
		if d.Name() == ".mjk" || d.Name() == ".mjkignore" || d.Name() == "." || d.Name() == "mjk" {
			return nil
		}
		// if directory make tree if not empty
		if d.IsDir() {
			// fmt.Println("Directory: ", d.Name())
			dirTree, empty := Tree(path, indexData)
			if empty {
				return nil
			}
			tree += "040000" + " " + d.Name() + " " + "\000" + dirTree
			isEmpty = false
			// fmt.Println(tree)
		} else {
			// check if the file exist in index file
			if fileData, ok := indexData.Files[path]; !ok {
				return nil
			} else {
				isEmpty = false
				tree += " " + fileData.FileType + " " + fileData.Name + "\000" + " " + fileData.Hash
			}
			// fmt.Println("File: ", d.Name())

			// fmt.Println(tree)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error in walking the directory")
	}
	if isEmpty {
		return "", isEmpty
	}
	tree = fmt.Sprintf("tree %d \000%s", len(tree), tree)
	// fmt.Println("final tree:" + tree)
	object_sha1 := utils.Sha1(tree)
	Object(object_sha1, fileUtils.Compress(tree))
	// fmt.Println("sha of tree: " + object_sha1)
	return object_sha1, isEmpty
}

// try os.ReadDir if more efficient
