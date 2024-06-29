package creator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func ReadIndex() modles.IndexFile {
	index, err := os.OpenFile(filepath.Join(".", ".mjk", "index"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error in reading index the file %v", err)
		os.Exit(1)
	}
	defer index.Close()
	indexInfo, _ := index.Stat()
	indexSize := indexInfo.Size()
	if indexSize == 0 {
		return modles.IndexFile{Header: modles.Header{NumberOfFiles: 0}, Files: make(map[string]modles.StageFile)}
	}
	// read the index file if not empty
	var indexData modles.IndexFile
	buffer := make([]byte, indexSize)
	index.Read(buffer)
	indexData.Decompress(*bytes.NewBuffer(buffer))
	return indexData
}

func ModifyIndexData(fileData *modles.StageFile, indexData *modles.IndexFile) {
	originalFileData, exists := indexData.Files[fileData.Path]
	// if new file, key in map do not exists
	if !exists {
		fileData.Hash = Blob(fileData.Path)
		indexData.Files[fileData.Path] = *fileData
		return
	}
	// if the file is modified by checnnking the mod time and the size
	if originalFileData.ModTime != fileData.ModTime || originalFileData.Size != fileData.Size {
		fileData.Hash = Blob(fileData.Path)
		return
	}
	// if the modifaction time and the size is same then the file is not modified
	// i am not sure
	fileData.Hash = originalFileData.Hash
}
func WriteIndex(indexData modles.IndexFile, rootPath string) {
	index, err := os.OpenFile(filepath.Join(rootPath, ".mjk", "index"), os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error in writing index the file %v", err)
		os.Exit(1)
	}
	defer index.Close()

	// check if the file is deleted
	for i, file := range indexData.Files {
		// file is now deleted
		if _, err := os.Stat(file.Path); err != nil {
			dir := filepath.Join(rootPath, ".mjk", "objects", file.Hash[:2])
			os.Remove(filepath.Join(dir, file.Hash[2:]))
			delete(indexData.Files, i)
			// delete the dir also if empty
			entries, _ := os.ReadDir(dir)
			if len(entries) == 0 {
				os.Remove(dir)
			}
		}
	}
	// write the remaining to the index file
	indexData.Header.NumberOfFiles = uint32(len(indexData.Files))
	compressedData := fileUtils.Compress(indexData)
	index.Write(compressedData.Bytes())
}
