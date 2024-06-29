package modles

import (
	"bytes"
	"os"
	"time"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

type StageFile struct {
	Path     string
	Name     string
	Size     int64
	ModTime  time.Time
	Hash     string
	FileType string
}

type Header struct {
	NumberOfFiles uint32
}
type IndexFile struct {
	Header Header
	Files  map[string]StageFile
}

func (data *IndexFile) Decompress(compressedData bytes.Buffer) {
	fileUtils.Decompress(compressedData, data)
}

func FileModeToGitMode(mode os.FileMode) string {
	if mode.IsDir() {
		return "040000" // Directory mode
	}
	if mode&0111 != 0 { // Check if any of the executable bits are set
		return "100755" // Executable file mode
	}
	// Regular file with 644 permissions
	return "100644"
}

func StageFileData(path string) StageFile {
	fileInfo, _ := os.Stat(path)
	fileData := StageFile{
		Path:     path,
		Name:     fileInfo.Name(),
		Hash:     "", // will add later when we check if the file is modtified or not
		ModTime:  fileInfo.ModTime(),
		Size:     fileInfo.Size(),
		FileType: FileModeToGitMode(fileInfo.Mode()),
	}
	return fileData
}
