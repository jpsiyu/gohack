package ch4

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type badAuth struct {
	Username string
	Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != b.Username && password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return
	}
	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)
	next(w, r)
}

func sayhello2(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "Hi %s\n", username)
}

func NegroniDemo() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayhello2).Methods("GET")

	n := negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "123456",
	})

	n.UseHandler(r)
	http.ListenAndServe(":8080", n)
}
