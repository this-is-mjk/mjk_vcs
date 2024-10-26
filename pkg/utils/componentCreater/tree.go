package creator

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils"
	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Tree(dirPath string, indexData modles.IndexFile, root string) (modles.Tree, bool) {
	isEmpty := true
	var tree modles.Tree
	tree.Signature = "tree"
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
			subTree, empty := Tree(path, indexData, root)
			if empty {
				return nil
			}
			tree.NumberOfEntries++
			tree.Entries = append(tree.Entries, modles.StageFile{FileType: "040000", Name: d.Name(), Hash: subTree.Hash})
			isEmpty = false
			// fmt.Println(tree)
		} else {
			// check if the file exist in index file
			if fileData, ok := indexData.Files[strings.TrimPrefix(path, root)[1:]]; !ok {
				return nil
			} else {
				// fmt.Println("File is in index: ", d.Name())
				isEmpty = false
				tree.NumberOfEntries++
				tree.Entries = append(tree.Entries, fileData)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error in walking the directory")
	}
	if isEmpty {
		return modles.Tree{}, isEmpty
	}
	tree.Hash = utils.Sha1(modles.ShowTree(tree))
	Object(tree.Hash, fileUtils.Compress(tree))
	return tree, false
}

// try os.ReadDir if more efficient
