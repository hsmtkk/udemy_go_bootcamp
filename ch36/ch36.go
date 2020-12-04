package main

import (
	"log"

	"github.com/hsmtkk/udemy_go_bootcamp/ch36/gover"
)

func main() {
	v := gover.Version()
	log.Print(v)
}
