package main

import (
	"fmt"
	"github.com/dablelv/go-huge-util/conv"
	"github.com/dablelv/go-huge-util/slice"
	"regexp"
)

// 4位数的数字-2020
func pruneYearSpec(yearSpec string) string {
	reg := regexp.MustCompile("\\d{4,}")
	strs := reg.FindAllString(yearSpec, -1)
	fmt.Println("strs: ", strs)

	ints, _ := conv.ToIntSliceE(strs)
	for i := 0; i < len(ints); i++ {
		ints[i] -= 2022
	}
	fmt.Println("ints: ", ints)

	prunedYearSpec, _ := slice.JoinWithSepE(ints, ",")
	return prunedYearSpec
}

func main() {
	yearSpec := "2026,2026,202612"
	prunedYearSpec := pruneYearSpec(yearSpec)
	fmt.Println("prunedYearSpec: ", prunedYearSpec)
}
