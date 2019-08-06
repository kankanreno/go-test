package main

import "fmt"

type stu struct {
	ID int
}

func foo(i int) func() int {
	return func() int {
		i++
		return i
	}
}


func main() {
	c1 := foo(0)
	c2 := foo(0)

	fmt.Printf("c1 type: %T\n", c1)
	fmt.Println("c1", c1)

	fmt.Printf("c2 type: %T\n", c2)
	fmt.Println("c2", c2)

	aaa := c1()
	fmt.Println("aaa = ", aaa)

	bbb := c2()
	fmt.Println("bbb = ", bbb)
}