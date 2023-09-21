package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"runtime"
)

//func main() {
//	numCPUs := runtime.NumCPU()
//	pool := tunny.NewFunc(numCPUs, func(payload any) any {
//		var result []byte
//
//		// TODO: Something CPU heavy with payload
//		str := strings.ToUpper(cast.ToString(payload))
//		time.Sleep(200 * time.Millisecond)
//		result = []byte(str)
//
//		return result
//	})
//	defer pool.Close()
//
//	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
//		input, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, "Internal error", http.StatusInternalServerError)
//		}
//		defer r.Body.Close()
//
//		// Funnel this work into our pool. This call is synchronous and will
//		// block until the job is completed.
//		result := pool.Process(input)
//
//		w.Write(result.([]byte))
//	})
//
//	http.ListenAndServe(":8080", nil)
//}

func main() {
	numCPUs := runtime.NumCPU()
	pool := tunny.NewCallback(numCPUs)

	printHello := func(str any) any {
		fmt.Println("Hello!", str)
		return "Hello! " + str.(string)
	}
	pool.Process(printHello("world"))
}

//type myWorker struct {
//	processor func(any) any
//}
//
//func (w *myWorker) Process(payload any) any {
//	return w.processor(payload)
//}
//
//func (w *myWorker) BlockUntilReady() {}
//func (w *myWorker) Interrupt()       {}
//func (w *myWorker) Terminate()       {}
//
//func main() {
//	printHello := func(str any) any {
//		fmt.Println("Hello!", str)
//		return "Hello! " + str.(string)
//	}
//
//	pool1 := tunny.New(3, func() tunny.Worker {
//		return &myWorker{
//			processor: printHello,
//		}
//	})
//	pool1.Process("world")
//}
