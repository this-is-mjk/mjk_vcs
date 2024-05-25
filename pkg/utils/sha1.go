package utils

import "crypto/sha1"

func Sha1(data string) []byte {
	// returns sha1 hash of the data
	var sha = sha1.New()
	sha.Write([]byte(data))
	// sha1 hash is 40 characters long
	return sha.Sum(nil)
}
