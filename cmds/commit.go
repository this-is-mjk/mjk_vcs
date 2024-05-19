package cmds

import (
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"mjk_vcs/helper"
	"os"
	"path/filepath"
	"strings"
)

const FILEMODE = 100644

func addFile(path string) []byte {
	data := helper.Read_file(path)
	// blob (length_of_content)(NULL_BYTE)(content)
	object_structure := fmt.Sprintf("blob %d\000%s", len(data), data)
	file_name := helper.Get_sha1(object_structure)
	helper.Write_file(hex.EncodeToString(file_name), helper.Compress(data))
	return file_name
}
func Commit() {
	var tree = ""
	err := filepath.WalkDir(helper.Get_pwd(), func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".mjk") {
			// ignores .mjk and .mjkignore
			return nil
		}
		// Currently things of only files
		if !d.IsDir() {
			fmt.Println(path, d.Name())
			// tree (length_of_content)(NULL_BYTE)(file1_MODE) (file1_NAME)(NULL_BYTE)(objID1)(file2_MODE) (file2_NAME)(NULL_BYTE)(objID2)(file3_NAME)(NULL_BYTE)..........
			tree += fmt.Sprintf("%d", FILEMODE) + " " + d.Name() + "\000" + string(addFile(path))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories\n Commit Cancled: %s", err)
	}
	tree = fmt.Sprint("tree ", len(tree), "\000", tree)
	tree_id := helper.Get_sha1(tree)
	helper.Write_file(hex.EncodeToString(tree_id), helper.Compress([]byte(tree)))
	// fmt.Print(string(tree_id), compress([]byte(tree)))

	// commit (length_of_content)(NULL_BYTE)tree (TREE_ID)
	// parent (PARENT_ID)(NULL_BYTE)
	// author (AUTHOR_NAME) (EMIAL) (TIME_STRING)
	// committer (COMMITTER_NAME) (EMAIL) (TIME_STRING)
	// (COMMIT_MESSAGE)

	// set environment in your shell by
	// Example: export AUTHOR_NAME=your_name
	commit := fmt.Sprintf("tree %s\n", hex.EncodeToString(tree_id))
	if !helper.Is_first_commit() {
		commit += fmt.Sprintf("parent %s\n", helper.Read_file(filepath.Join(helper.Get_pwd(), ".mjk", "HEAD")))
	}
	commit += fmt.Sprintf("author %s %s\ncommitter %s %s\n", os.Getenv("AUTHOR_NAME"), os.Getenv("AUTHOR_EMAIL"), os.Getenv("COMMITTER_NAME"), os.Getenv("COMMITTER_EMAIL"))
	commit = fmt.Sprintf("commit %d\000 %s", len(commit), commit)
	fmt.Println(commit)
	commit_id := helper.Get_sha1(commit)
	helper.Write_file(hex.EncodeToString(commit_id), helper.Compress([]byte(commit)))
	// write HEAD
	file_dir := filepath.Join(helper.Get_pwd(), ".mjk", "HEAD")
	f, err := os.Create(filepath.Join(file_dir))
	if err != nil {
		log.Fatalf("Unable to write file\nERROR:- %s", err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(hex.EncodeToString(commit_id))); err != nil {
		log.Fatalf("Error writeing compressed_data\nERROR:- %s", err)
	}

}
