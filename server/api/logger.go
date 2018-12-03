package api

import (
	"log"
	"os"
)

// NewLogger ...
func NewLogger(file *os.File) *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	//return log.New(file, "", log.Ldate|log.Ltime)
}
