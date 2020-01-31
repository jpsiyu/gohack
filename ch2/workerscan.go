package ch2

import (
	"fmt"
	"net"
	"strconv"
)

func worker(ports, results chan int) {
	host := "scanme.nmap.org"
	for p := range ports {
		address := host + ":" + strconv.Itoa(p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func WorkerScan() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	maxPort := 1024

	defer close(ports)
	defer close(results)

	// start workers
	// every time a port was pushed to ports buffer, multiple workers compete for it,
	// and only one winner will get the job.
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// fill ports
	go func() {
		for i := 0; i < maxPort; i++ {
			ports <- i
		}
	}()

	// read results
	for i := 0; i < maxPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// print the open ports
	fmt.Println(openports)
}
