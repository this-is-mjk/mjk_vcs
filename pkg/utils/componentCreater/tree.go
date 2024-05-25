package creator

import (
	"encoding/hex"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Tree(dirPath string) string {
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
		// if directory form tree
		if d.IsDir() {
			// fmt.Println("Directory: ", d.Name())
			tree += "040000" + " " + d.Name() + "\000" + string(Tree(path))
			// fmt.Println(tree)
		} else {
			// fmt.Println("File: ", d.Name())
			tree += "100644" + " " + d.Name() + "\000" + " " + string(Blob(path))
			// fmt.Println(tree)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error in walking the directory")
	}
	tree = fmt.Sprintf("tree %d \000%s", len(tree), tree)
	// fmt.Println("final tree: " + tree)
	object_sha1 := hex.EncodeToString(utils.Sha1(tree))
	Object(object_sha1, fileUtils.Compress(tree))
	// fmt.Println("sha of tree: " + object_sha1)
	return object_sha1
}

// try os.ReadDir if more efficient