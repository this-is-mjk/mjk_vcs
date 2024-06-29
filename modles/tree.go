package modles

import (
	"bytes"
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

type Tree struct {
	Signature       string
	NumberOfEntries int
	Entries         map[string]StageFile
}

func (data *Tree) Decompress(compressedData bytes.Buffer) {
	fileUtils.Decompress(compressedData, data)
}
func ReadTree(Id string) Tree {
	data, _ := fileUtils.ReadFile(filepath.Join(".", ".mjk", "objects", Id[:2], Id[2:]))
	var tree Tree
	tree.Decompress(*bytes.NewBuffer(data))
	return tree
}
func ShowTree(tree Tree) string {
	returnString := ""
	for key, value := range tree.Entries {
		returnString += key + " " + value.FileType + "\n"
	}
	return returnString
}
