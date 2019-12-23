package main

import "fmt"

func incr() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	f := incr()

	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
}