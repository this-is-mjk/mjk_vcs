package modles

import (
	"bytes"

	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

type Blob struct {
	Signature string
	Size      int
	Content   string
}

func (data *Blob) Decompress(compressedData bytes.Buffer) {
	fileUtils.Decompress(compressedData, data)
}

func CreateBlob(content string) Blob {
	return Blob{
		Signature: "blob",
		Size:      len(content),
		Content:   content,
	}
}
