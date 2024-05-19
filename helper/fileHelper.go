package helper

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"log"
	"os"
	"path/filepath"
)

func Read_file(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file\nERROR:- %s", err)
	}
	return data
}
func Compress(data []byte) bytes.Buffer {
	var buffer bytes.Buffer
	w := zlib.NewWriter(&buffer)
	w.Write(data)
	w.Close()
	return buffer
}
func Get_sha1(data string) []byte {
	var sha = sha1.New()
	sha.Write([]byte(data))
	return sha.Sum(nil)
}
func Write_file(file_name string, compressed_data bytes.Buffer) {
	file_dir := filepath.Join(Get_pwd(), ".mjk", "objects", file_name[:2])
	CreateDir(file_dir)
	f, err := os.Create(filepath.Join(file_dir, file_name[2:]))
	if err != nil {
		log.Fatalf("Unable to write file1\nERROR:- %s", err)
	}
	defer f.Close()
	if _, err := compressed_data.WriteTo(f); err != nil {
		log.Fatalf("Error writeing compressed_data\nERROR:- %s", err)
	}
}
