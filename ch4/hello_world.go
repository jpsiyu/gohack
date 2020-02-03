package ch4

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func HelloWorldServe() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
