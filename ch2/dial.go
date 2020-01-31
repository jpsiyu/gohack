package ch2

import (
	"fmt"
	"net"
	"strconv"
)

func Dial() {
	for i := 1; i <= 1024; i++ {
		address := "scanme.nmap.org:" + strconv.Itoa(i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("port %d is open\n", i)
	}
}
