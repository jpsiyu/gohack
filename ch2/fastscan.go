package ch2

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func FastScan() {
	host := "scanme.nmap.org"
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := host + ":" + strconv.Itoa(port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("port %d is open\n", port)
		}(i)
	}
	wg.Wait()
	fmt.Println("scan finish")
}
