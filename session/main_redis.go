package main

import (
	"fmt"
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
)

var store, _ = redistore.NewRediStoreWithDB(10, "tcp", "47.94.212.166:6379", "Cangbai.123", "14", []byte("secret-key"))

func set(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_id")

	session.Values["name"] = "看看"
	session.Values["account"] = "kankan"

	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hello World")
}

func read(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_id")
	fmt.Fprintf(w, "name: %s, account: %s", session.Values["name"], session.Values["account"])
}

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
