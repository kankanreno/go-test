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

	str := ""
	for i := 0; i < 40; i++ {
		go func(str string) {
			for j := 0; j < 100000; j++ {
				str += "0123456789"
			}
		}(str)
	}

	time.Sleep(10000 * time.Hour)
}
