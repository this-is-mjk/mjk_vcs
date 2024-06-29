package creator

import (
	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils"

	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Blob(path string) string {
	// Read the file and form blob
	data, _ := fileUtils.ReadFile(path)
	blob := modles.CreateBlob(string(data))
	// fmt.Println(blob)
	blob_sha1 := utils.Sha1(blob)
	// create object
	Object(blob_sha1, fileUtils.Compress(blob))
	return blob_sha1
}
