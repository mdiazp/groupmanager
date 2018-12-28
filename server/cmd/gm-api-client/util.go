package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

func readString() string {
	s, e := rd.ReadString('\n')
	if e != nil {
		panic(e)
	}
	return s[:len(s)-1]
}

func readInt() int {
	s := readString()
	num, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return num
}

// WR ...
func WR(resp *http.Response, body []byte) {
	fmt.Printf("StatusCode: %d Status: %s\n", resp.StatusCode, resp.Status)
	fmt.Println(string(body))
}

func pe(e error) {
	if e == nil {
		return
	}
	log.Fatalf("Location: %s\n\tError: %s\n", WAI(2), e.Error())
	panic(e)
}

// WAI ...
func WAI(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s:%d", file, line)
}
