package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//wg.Add(1)
	//timer := time.NewTimer(time.Second * 1)
	//defer timer.Stop()
	//go func() {
	//	defer wg.Done()
	//
	//	for {
	//		<-timer.C
	//		fmt.Println("timer.")
	//		timer.Reset(time.Second * 1)
	//	}
	//}()
	//
	//wg.Add(1)
	//ticker := time.NewTicker(time.Second * 1)
	//defer ticker.Stop()
	//go func() {
	//	wg.Done()
	//
	//	for {
	//		<-ticker.C
	//		fmt.Println("ticker.")
	//	}
	//}()
	//
	//wg.Wait()

	syncTicker := time.NewTicker(time.Second)
	defer syncTicker.Stop()
	housekeepingTicker := time.NewTicker(2 * time.Second)
	defer housekeepingTicker.Stop()

	for {
		syncLoopIteration(syncTicker.C, housekeepingTicker.C)
	}
}

func syncLoopIteration(syncCh <-chan time.Time, housekeepingCh <-chan time.Time) {
	select {
	case <-syncCh:
		fmt.Println("syncCh.")
	case <-housekeepingCh:
		fmt.Println("housekeepingCh.")
	}
}
