package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	//http.HandleFunc("/", handlerData)
	//http.ListenAndServe(":8080", nil)
	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}

//func handlerData(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "hello world")
//	//w.Write([]byte("Hello world!\n"))
//}