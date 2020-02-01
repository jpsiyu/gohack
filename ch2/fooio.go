package ch2

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")

	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func IODemo() {
	var (
		reader FooReader
		writer FooWriter
	)
	input := make([]byte, 4096)
	lenght, err := reader.Read(input)
	if err != nil {
		log.Println("unable to read data")
	}
	fmt.Printf("read %d bytes from stdin\n", lenght)

	lenght, err = writer.Write(input)
	if err != nil {
		log.Println("unable to write data")
	}
	fmt.Printf("write %d bytes to stdout\n", lenght)
}
