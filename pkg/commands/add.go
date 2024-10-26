package commands

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/this-is-mjk/mjk/modles"
	creator "github.com/this-is-mjk/mjk/pkg/utils/componentCreater"
)

func Add(filesToAdd []string) {
	if len(filesToAdd) == 0 {
		fmt.Println("No files to add, please provide")
		return
	}
	// // read index file
	// pwd := fileUtils.Get_pwd()
	pwd := "."

	// Read index file
	IndexFile := creator.ReadIndex()

	// sort the files to add lexographically
	sort.Strings(filesToAdd)

	for _, file := range filesToAdd {
		err := filepath.WalkDir(filepath.Join(pwd, file), func(path string, d fs.DirEntry, err error) error {
			// if file is .mjk or .mjkignore file return back
			if strings.Contains(path, ".mjk") {
				return nil
			}
			// if directory return back
			if d.IsDir() {
				return nil
			}
			// make object and the meta data
			fileData := modles.StageFileData(path)
			creator.ModifyIndexData(&fileData, &IndexFile)
			return nil
		})
		if err != nil {
			fmt.Println("Error in staging files")
		}
	}
	creator.WriteIndex(IndexFile, pwd)

	// creator.StageFile(pwd)

}
