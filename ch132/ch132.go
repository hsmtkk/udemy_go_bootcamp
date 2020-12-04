package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/udemy_go_bootcamp/ch132/valid"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s username password", os.Args[0])
	}
	username := os.Args[1]
	password := os.Args[2]
	if valid.New().Validate(username, password) {
		fmt.Println("accepted")
	} else {
		fmt.Println("wrong username or password")
	}
}
