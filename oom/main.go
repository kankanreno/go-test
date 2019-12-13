package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":9999", nil))
	}()

	//http.HandleFunc("/ping", handlerData)
	//http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlerData)
	http.ListenAndServe(":8080", mux)
}

func handlerData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	//w.Write([]byte("Hello world!\n"))
}