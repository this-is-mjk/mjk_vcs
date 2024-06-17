package creator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/this-is-mjk/mjk/structs"
)

func fileModeToGitMode(mode os.FileMode) string {
	if mode.IsDir() {
		return "040000" // Directory mode
	}
	if mode&0111 != 0 { // Check if any of the executable bits are set
		return "100755" // Executable file mode
	}
	// Regular file with 644 permissions
	return "100644"
}

func writeIndexFile(dirPath string, stagingFile structs.StageFile) {
	toWrite := fmt.Sprintf("%s\000%d\000%s\000%s\000%s\000%s\000", stagingFile.FileType, stagingFile.Size, stagingFile.Name, stagingFile.Time, stagingFile.Hash, stagingFile.DirPath)
	toWrite = fmt.Sprintf("%d\000%s\n", len(toWrite), toWrite)
	file, _ := os.OpenFile(dirPath+"/.mjk/index", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(toWrite)

}
func StageFile(dirPath string) {
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		// if file is .mjk or .mjkignore file return back
		if strings.Contains(path, ".mjk") {
			return nil
		}
		// if directory return back
		if d.IsDir() {
			return nil
		} else {
			fileInfo, _ := os.Stat(path)
			fileData := structs.StageFile{
				DirPath:  strings.Replace(path, dirPath+"/", "", 1),
				Name:     d.Name(),
				Hash:     string(Blob(path)),
				Time:     fileInfo.ModTime(),
				Size:     fileInfo.Size(),
				FileType: fileModeToGitMode(fileInfo.Mode()),
			}
			writeIndexFile(dirPath, fileData)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error in walking the directory")
	}
}
// // other method to get all files and directories in the directory
	// files, err1 := os.ReadDir(dirPath)
	// if err1 != nil {
	// 	fmt.Println("Error in reading the directory")

	// }
	// for f := range files {
	// 	// if file is .mjk or .mjkignore file ignore it
	// 	if strings.Contains(files[f].Name(), ".mjk") {
	// 		continue
	// 	}
	// 	// if directory call again on it
	// 	fmt.Println(files[f].IsDir())
	// 	fmt.Println(files[f].Type())

	// }
