package modles

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

type Tree struct {
	Signature       string
	Hash            string
	NumberOfEntries int
	Entries         []StageFile
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
	returnString := tree.Signature + fmt.Sprintf(" %d\n", tree.NumberOfEntries)
	for _, entry := range tree.Entries {
		returnString += entry.FileType + " " + entry.Name + " " + entry.Hash + "\n"
	}
	return returnString
}
