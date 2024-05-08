package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	//var id uint = 2
	//r := lo.Filter([]string{"1", "2", "3", "4"}, func(x string, _ int) bool {
	//	return x != cast.ToString(id)
	//})
	//fmt.Println(r)

	//r := lo.Intersect([]string{"1", "2", "3", "4"}, []string{"2", "3", "7"})
	//fmt.Println(r)

	ok := lo.Some([]string{"1", "2", "3", "4"}, []string{"2", "3", "7"})
	fmt.Println(ok)
}
