package api

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// NewLogger ...
func NewLogger(file *os.File) *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// return log.New(file, "", log.Ldate|log.Ltime)
}

// WAI ...
func WAI(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s:%d", file, line)
}
