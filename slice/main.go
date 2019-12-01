package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
)

func main() {
	var id uint = 2
	r := funk.Filter([]string{"1", "2", "3", "4"}, func(x string) bool {
		return x != cast.ToString(id)
	})
	fmt.Println(r)
	for _, v := range r.([]string) {
		fmt.Println(v)
	}
}
