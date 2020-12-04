package main

import (
	"fmt"

	"github.com/hsmtkk/udemy_go_bootcamp/ch157/partsofday"
)

func main() {
	g := partsofday.New().GetCurrent()
	fmt.Println(g)
}
