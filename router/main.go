package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := NewRouter()
	r.HandleFunc("GET", "/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "FOOOOOOOO!")
	})

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ROOT!")
	})

	http.ListenAndServe(":8080", r)
}
