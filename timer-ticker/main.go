package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	timer := time.NewTimer(time.Second * 1)
	defer timer.Stop()
	go func() {
		defer wg.Done()

		for {
			<-timer.C
			fmt.Println("timer.")
			timer.Reset(time.Second * 1)
		}
	}()

	wg.Add(1)
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	go func() {
		wg.Done()

		for {
			<-ticker.C
			fmt.Println("ticker.")
		}
	}()

	wg.Wait()
}
