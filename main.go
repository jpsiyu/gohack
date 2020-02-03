package main

import (
	"fmt"

	"log"

	"github.com/jpsiyu/gohack/ch4"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(("running..."))
	ch4.NegroniDemo()
}
