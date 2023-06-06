package main

import (
	"fmt"
)

func main() {

	////// 随机性测试
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	foo := 0

	for {
		select {
		case foo = <-ch1:
			fmt.Println("ch1 get success: ", foo)
		case foo = <-ch2:
			fmt.Println("ch2 get success: ", foo)
		default:
			fmt.Println("default")
			return
		}
	}

	////// 超时判断
	//ch := make(chan int)
	//timeout := make(chan bool, 1)
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	timeout <- true
	//}()
	//
	//select {
	//case <-ch:
	//case <-timeout:
	//	fmt.Println("timeout!")
	//case <-time.After(500 * time.Millisecond):
	//	fmt.Println("500 Millisecond timeout!")
	//}

	////// 判断 channel 是否是为空
	//ch := make(chan int, 5)
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//ch <- 4
	////ch <- 5
	//select {
	//case ch <- 100:
	//	fmt.Println("add success")
	//default:
	//	fmt.Println("channel 满了")
	//}
	//
	//close(ch)
	//
	//for v := range ch {
	//	fmt.Println(v)
	//}

	////// 退出
	//shouldQuit := make(chan struct{})
	//
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	close(shouldQuit)
	//}()
	//
	//select {
	//case <-shouldQuit:
	//	fmt.Println("quited!")
	//}
}
