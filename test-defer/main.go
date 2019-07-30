package main

import "fmt"

func main() {
	fmt.Printf("有名返回 = %d\n", test1())
	fmt.Printf("匿名返回 = %d\n", test2())
}

// 有名返回
func test1() (foo int) {
	defer func () {
		foo++
	}()

	return
}

// 匿名返回
func test2() int {
    foo := 0

    defer func () {
        foo++
        fmt.Printf("匿名返回 defer 中，foo = %d\n", foo)
    }()

	return foo
}
