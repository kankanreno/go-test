package main

import "fmt"

func main() {
	fmt.Printf("test1 = %d\n", test1())
	fmt.Printf("test2 = %d\n", test2())
}

// 匿名返回
func test1() int {
    foo := 0

    defer func () {
        foo++
        fmt.Printf("foo = %d\n", foo)
    }()

	return foo
}

// 有名返回
func test2() (foo int) {
    defer func () {
        foo++
    }()

	return
}
