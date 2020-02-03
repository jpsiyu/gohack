package ch4

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("start %s\n", time.Now().String())
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func MiddlewareDemo() {
	f := http.HandlerFunc(sayhello)
	l := logger{Inner: f}
	http.ListenAndServe(":8080", &l)
}
