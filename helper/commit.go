package helper

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const FILEMODE = 100644

func CreateDir(name string) {
	if err := os.MkdirAll(name, os.ModePerm); err != nil {
		log.Fatalf("Unable to create directory\nERROR:- %s", err)
	}
}
func get_pwd() string {
	// returns working dir in string
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to fetch present working directory\n ERROR: %s", err)
	}
	return wd
}
func read_file(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file\nERROR:- %s", err)
	}
	return data
}
func compress(data []byte) bytes.Buffer {
	var buffer bytes.Buffer
	w := zlib.NewWriter(&buffer)
	w.Write(data)
	w.Close()
	return buffer
}
func get_sha1(data string) []byte {
	var sha = sha1.New()
	sha.Write([]byte(data))
	return sha.Sum(nil)
}
func write_file(file_name string, compressed_data bytes.Buffer) {
	file_dir := filepath.Join(get_pwd(), ".mjk", "objects", file_name[:2])
	CreateDir(file_dir)
	f, err := os.Create(filepath.Join(file_dir, file_name[2:]))
	if err != nil {
		log.Fatalf("Unable to write file1\nERROR:- %s", err)
	}
	defer f.Close()
	if _, err := compressed_data.WriteTo(f); err != nil {
		log.Fatalf("Error writeing compressed_data\nERROR:- %s", err)
	}
}
func addFile(path string) []byte {
	data := read_file(path)
	// blob (length_of_content)(NULL_BYTE)(content)
	object_structure := fmt.Sprintf("blob %d\000%s", len(data), data)
	file_name := get_sha1(object_structure)
	write_file(hex.EncodeToString(file_name), compress(data))
	return file_name
}
func Commit() {
	var tree = ""
	err := filepath.WalkDir(get_pwd(), func(path string, d fs.DirEntry, err error) error {
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
	tree += fmt.Sprint("tree ", len(tree), "\000", tree)
	tree_id := get_sha1(tree)
	write_file(hex.EncodeToString(tree_id), compress([]byte(tree)))
	// fmt.Print(string(tree_id), compress([]byte(tree)))

	// commit (length_of_content)(NULL_BYTE)tree (TREE_ID)
	// parent (PARENT_ID)(NULL_BYTE)author
	// author (AUTHOR_NAME) (EMIAL) (TIME_STRING)
	// committer (COMMITTER_NAME) (EMAIL) (TIME_STRING)
	// (COMMIT_MESSAGE)
	commit := fmt.Sprintf("tree %d\nauthor %s %s\ncommitter %s %s\n", tree_id, os.Getenv("AUTHOR_NAME"), os.Getenv("AUTHOR_EMAIL"), os.Getenv("COMMITTER_NAME"), os.Getenv("COMMITTER_EMAIL"))
	commit += fmt.Sprintf("commit %d\000 %s", len(commit), commit)
	commit_id := get_sha1(commit)
	write_file(hex.EncodeToString(commit_id), compress([]byte(commit)))
	// write HEAD
	file_dir := filepath.Join(get_pwd(), ".mjk", "HEAD")
	f, err := os.Create(filepath.Join(file_dir))
	if err != nil {
		log.Fatalf("Unable to write file\nERROR:- %s", err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(hex.EncodeToString(commit_id))); err != nil {
		log.Fatalf("Error writeing compressed_data\nERROR:- %s", err)
	}

}
