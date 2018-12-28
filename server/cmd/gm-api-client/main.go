package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	apiURL = "http://localhost:1234/"
	ct     = "application/json; charset=UTF-8"
)

var rd *bufio.Reader

func main() {
	rd = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Controller Name: ")
		op := readString()
		switch op {
		case "account.register":
			Register()
		case "user.retrieve":
			userRetrieve()
		case "user.update":
			userUpdate()
		}
	}
}
