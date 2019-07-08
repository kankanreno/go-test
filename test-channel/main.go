package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
	for i := 0; i < 10; i++ {
		e, ok := <-ch
		fmt.Printf("%v, %v\n", e, ok)

		if !ok {
			break
		}
	}
}
