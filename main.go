package main

import (
	"fmt"

	"log"

	"github.com/jpsiyu/gohack/ch3"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(("running..."))
	ch3.Demo35()
}
