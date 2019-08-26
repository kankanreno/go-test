package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//func foo(ch chan int) {
//	defer wg.Done()
//	fmt.Println("in goroutine")
//
//	for i := 0; i < 10; i++ {
//		//e, ok := <-ch
//		//fmt.Printf("%v, %v\n", e, ok)
//		//
//		//if !ok {
//		//	break
//		//}
//		e := <-ch
//		fmt.Printf("%v\n", e)
//	}
//}
//
//func main() {
//	ch := make(chan int, 5)
//
//	for i := 0; i < 5; i++ {
//		ch <- i
//	}
//
//	close(ch)
//
//	wg.Add(1)
//	go foo(ch)
//
//	wg.Wait()
//}

func main() {
	foo := make(chan int)

	for {
		go func() {
			time.Sleep(1 * time.Second)
			foo <- 0
		}()

		fmt.Println("for")
		select {
		case <-foo:
			fmt.Println("hha!")
		//default:
		//	fmt.Println("default!")
		}
	}
}
