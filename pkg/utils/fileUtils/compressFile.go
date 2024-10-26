package fileUtils

import (
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
)

func Compress(data interface{}) bytes.Buffer {
	var bufferForGob bytes.Buffer
	// Create a new gob encoder and encode the commit struct
	encoder := gob.NewEncoder(&bufferForGob)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding Gob:", err)
		os.Exit(1)
	}
	// compress file
	var bufferForZlib bytes.Buffer
	w := zlib.NewWriter(&bufferForZlib)
	w.Write(bufferForGob.Bytes())
	w.Close()
	return bufferForZlib
}

func Decompress(compressedData bytes.Buffer, data interface{}) {
	r, err := zlib.NewReader(&compressedData)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	decompressedData, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// Create a buffer with the decompressed data
	bufferForGob := bytes.NewBuffer(decompressedData)
	fmt.Printf("Decompressed data: %s\n", bufferForGob)
	// Create a new gob decoder and decode the data into the provided struct
	decoder := gob.NewDecoder(bufferForGob)
	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("Error decoding Gob:", err)
		os.Exit(1)
	}
}
