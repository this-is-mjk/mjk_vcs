package creator

import (
	"bytes"
	"fmt"
	"path/filepath"

	fileUtils "github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Object(name string, data bytes.Buffer) {
	// create parent dir from first 2 chars
	fileUtils.CreateDirectory(filepath.Join("./.mjk/objects", name[:2]))
	// write into the file
	fileUtils.WriteFile(fmt.Sprintf("./.mjk/objects/%s/%s", name[:2], name[2:]), data)
}
