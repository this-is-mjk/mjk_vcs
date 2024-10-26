package commands

import (
	"fmt"

	"github.com/this-is-mjk/mjk/modles"
)

func Test() {
	fmt.Println("Test")
	// blob := modles.Blob{
	// 	Content: "abcdefghijklmnopqrstuvwxyz1234567890",
	// }
	// fmt.Println(blob)
	// compress := fileUtils.Compress(blob)
	// fmt.Print(compress)
	// var decompressedBlob modles.Blob
	// decompressedBlob.Decompress(compress)
	// fmt.Println(decompressedBlob)

	fmt.Println(modles.ReadCommit("236aab668efa5e070587285d6d3df7f7db542537"))
}
