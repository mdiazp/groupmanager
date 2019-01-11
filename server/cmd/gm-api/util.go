package main

import (
	"log"
	"os"
)

// FileToBytes ...
func FileToBytes(f *os.File) []byte {
	fileinfo, e := f.Stat()
	if e != nil {
		log.Fatalf("File at function FileToBytes")
		panic(e)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, e = f.Read(buffer)
	if e != nil {
		log.Fatalf("File at function FileToBytes")
		panic(e)
	}
	return buffer
}
