package main

import (
	"fmt"
	"time"
)

func main() {
	//var slice []int
	////slice := make([]int, 0)
	//
	//if slice == nil {
	//	fmt.Print("slice is nil\n")
	//}
	//
	//fmt.Printf("slice, %v\n", slice)
	//fmt.Printf("slice, %d\n", len(slice))
	foo := []string{"stanley", "david", "oscar"}
	var bar interface{}
	bar = foo
	fmt.Printf("foo, %+v\n", foo)
	fmt.Printf("bar, %+v\n", bar)

	var m []byte
	for i := 0; i < 15; i++ {
		m = append(m, make([]byte, 1024 * 1024 * 10)...)
		time.Sleep(1 * time.Second)
		fmt.Printf("loop: %d, sizeof(m): %dMB\n", i, len(m) / (1024 * 1024))
	}

	time.Sleep(10000 * time.Hour)
}
