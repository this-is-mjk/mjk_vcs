package fileUtils

import (
	"bytes"
	"compress/zlib"
)

func Compress(data string) bytes.Buffer {
	// compress file
	var buffer bytes.Buffer
	w := zlib.NewWriter(&buffer)
	w.Write([]byte(data))
	w.Close()
	return buffer
}