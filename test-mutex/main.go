package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex
var ch = make(chan bool, 1)
var x = 0

func increment() {
	defer wg.Done()
	mux.Lock()
	//ch <- true		// 由于缓冲信道的容量为 1，所以任何其它协程试图写入该信道时，都会发生阻塞
	x += 1
	//<-ch
	mux.Unlock()
}

func main() {
	for i := 0; i < 10000000; i++ {
		wg.Add(1)

		go increment()
	}

	wg.Wait()
	fmt.Println("final value of x:", x)
}
