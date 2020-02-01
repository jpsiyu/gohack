package ch2

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		log.Fatalln("Unable to connect dst site")
	}
	defer dst.Close()
	// forward src to dst
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// forward dst to src
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func ProxyDemo() {
	port := ":1990"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Unable to listen on port", port, err)
	}
	log.Println("server listening or port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
