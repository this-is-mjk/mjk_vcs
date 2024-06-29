package commands

import (
	"fmt"

	"github.com/this-is-mjk/mjk/modles"
	"github.com/this-is-mjk/mjk/pkg/utils/fileUtils"
)

func Test() {
	fmt.Println("Test")
	blob := modles.Blob{
		Content: "abcdefghijklmnopqrstuvwxyz1234567890",
	}
	fmt.Println(blob)
	compress := fileUtils.Compress(blob)
	fmt.Print(compress)
	var decompressedBlob modles.Blob
	decompressedBlob.Decompress(compress)
	fmt.Println(decompressedBlob)
}
