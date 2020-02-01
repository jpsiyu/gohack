package main

import (
	"fmt"

	"log"

	"github.com/jpsiyu/gohack/ch2"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fmt.Println(("running..."))
	ch2.ProxyDemo()
}
