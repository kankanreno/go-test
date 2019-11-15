package main

import (
	"fmt"
	"regexp"
)

func main() {
    //ok, _ := regexp.MatchString("^-?[0-9a-zA-Z_]$", "-hello123")
    ok, _ := regexp.MatchString("^-?\\w$", "-hello123")
	fmt.Println(ok)
}
