package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
)

func Sha1(data interface{}) string {
	var bufferForGob bytes.Buffer
	// Create a new gob encoder and encode the commit struct
	encoder := gob.NewEncoder(&bufferForGob)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding Gob:", err)
		os.Exit(1)
	}
	// returns sha1 hash of the data
	var sha = sha1.New()
	sha.Write(bufferForGob.Bytes())
	// sha1 hash is 40 characters long
	return hex.EncodeToString(sha.Sum(nil))
}
