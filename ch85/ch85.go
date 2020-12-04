package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/udemy_go_bootcamp/ch85/bang"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s arg", os.Args[0])
	}
	banger := bang.New()
	fmt.Println(banger.Bang(os.Args[1]))
}
