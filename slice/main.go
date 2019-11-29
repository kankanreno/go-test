package main

import (
	"fmt"
    "github.com/wesovilabs/koazee"
)

func main() {
    var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}
    fmt.Println(stream.Drop(5).Do().Out().Val())
}
