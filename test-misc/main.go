package main

import "fmt"

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
}
