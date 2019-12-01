package main

import (
	"fmt"
    "github.com/thoas/go-funk"
)

func main() {
	r := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x != 2
	}) // []int{2, 4}
    fmt.Println(r)
}
