package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str := "-hello123"
	////ok, _ := regexp.MatchString("^-?[0-9a-zA-Z_]$", str)
	//ok, _ := regexp.MatchString("^-?\\w$", str)
	//fmt.Println(ok)

	//str := ""
	//ok, _ := regexp.MatchString("^\\w+:\\d:.*;?$", str)
	//fmt.Println(ok)

	str := "QN000008"
	reg := regexp.MustCompile(`[1-9]+0*`)
	match := reg.FindAllString(str, -1)
	fmt.Println(match) // 输出：[000008]

	//num := cast.ToInt(match[0])
	//fmt.Println(num)

	//reg, err := regexp.Compile("(\\w*)(\\d+)")
	//
	//if err != nil {
	//	fmt.Printf("error compiling regexp: %v\n", err)
	//	return
	//}
	//
	//matchStrings := reg.FindStringSubmatch(str)
	//fmt.Println(matchStrings)
}
