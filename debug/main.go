package main

import "fmt"

func executeFn(fn func() int) int {
    return fn()
}

func main() {
    a := 1
    b := 2
    c := executeFn(func() int {
        a += b
        return a
    })
    fmt.Printf("%d %d %d\n", a, b, c)
}