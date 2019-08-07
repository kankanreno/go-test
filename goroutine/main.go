package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func printNumber(from, to int, c chan int) {
	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep( 1 * time.Millisecond)
	}

	c <- 0
}

//func main() {
//	//c := make(chan int, 3)
//	//go printNumber(1, 3, c)
//	//go printNumber(4, 6, c)
//	//_ = <- c
//	//_ = <- c
//	go func() {
//		fmt.Println("go1 begin...")
//		go func() {
//			fmt.Println("go2 begin...")
//			time.Sleep(1 * time.Second)
//			fmt.Println("go2 end")
//		}()
//		fmt.Println("go1 end")
//	}()
//
//	time.Sleep(3 * time.Second)
//}

var wg sync.WaitGroup
var mux sync.Mutex

type Data struct {
	ID    int
}

func foo(i int, data *Data) {
	defer wg.Done()

	logrus.Infof("data address: %p", data)
	mux.Lock()
	data.ID = i
	if data.ID != i {
		logrus.Info(i, data.ID)
	}
	mux.Unlock()
}

func main() {
	//data1 := &Data{}
	//logrus.Infof("data1 address: %p", data1)
	//data2 := &Data{}
	//logrus.Infof("data2 address: %p", data2)

	data := &Data{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go foo(i, data)
	}

	wg.Wait()
}