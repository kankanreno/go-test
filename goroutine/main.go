package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func add(i int) {
	defer wg.Done()

	x := 0
	for i := 0; i < 10000000000; i++ {
		x++
	}
	log.Printf("%d, %d", i, x)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go add(i)
	}

	wg.Wait()
}