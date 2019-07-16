package main

import (
	"github.com/Jeffail/tunny"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	numCPUs := runtime.NumCPU()
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte
		str := strings.ToUpper(cast.ToString(payload))
		result = []byte(str)
		return result
	})
	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		result := pool.Process(input)
		w.Write(result.([]byte))
	})
	http.ListenAndServe(":8080", nil)
}
