package creator

import (
	"encoding/hex"
	"fmt"

	"github.com/this-is-mjk/mjk/pkg/utils"

	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Blob(path string) string {
	// Read the file
	data := fileUtils.ReadFile(path)
	// Create a blob >> blob (length) + \0 + data
	blob := fmt.Sprintf("blob %d \000%s", len(data), data)
	// fmt.Println(blob)
	blob_sha1 := hex.EncodeToString(utils.Sha1(blob))
	Object(blob_sha1, fileUtils.Compress(blob))
	return blob_sha1
}
