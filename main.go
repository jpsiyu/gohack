package main

import (
	"fmt"

	"log"

	"github.com/jpsiyu/gohack/ch5"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(("running..."))
	ch5.FlagParser()
}
