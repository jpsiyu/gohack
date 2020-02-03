package ch4

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintf(w, "Executing /a\n")
	case "/b":
		fmt.Fprintf(w, "Executing /b\n")
	case "/c":
		fmt.Fprintf(w, "Executing /c\n")
	default:
		http.Error(w, "404 not found", 404)
	}
}

func SimpleRouter() {
	var r router
	http.ListenAndServe(":8080", &r)
}
