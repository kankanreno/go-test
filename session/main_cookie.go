package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

// REF: https://xie.infoq.cn/article/53d671122ad09f97080fdab35		「Go 工具箱」gorilla/sessions 包的使用及原理分析

// custom session stores
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

func set(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_id")
	session.Values["account"] = "kankan"
	session.Values["name"] = "看看"
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hello World")
}

func read(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_id")
	fmt.Fprintf(w, "account: %s, name: %s", session.Values["account"], session.Values["name"])
}

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
