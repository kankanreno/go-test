package main

import (
	"fmt"
)

func main() {
	column := "A"
	prevColumn := GetPrevColumn(column)
	fmt.Printf("prev column: %s -> %s\n", column, prevColumn)

	//column = "B"
	//nextColumn := GetNextColumn(column)
	//fmt.Printf("next column: %s -> %s\n", column, nextColumn)
}

func GetPrevColumn(column string) string {
	if column == "A" {
		return ""
	}

	prefix := column[0 : len(column)-1]
	last := column[len(column)-1]
	if last != 'A' {
		return prefix + string(rune(last-1))
	}

	return GetPrevColumn(prefix) + "Z"
}

func GetNextColumn(column string) string {
	if column == "" {
		return "A"
	}

	prefix := column[0 : len(column)-1]
	last := column[len(column)-1]
	if last != 'Z' {
		return prefix + string(rune(last+1))
	}

	return GetNextColumn(prefix) + "A"
}
