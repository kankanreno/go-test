package main

import (
	"fmt"
	"regexp"
)

func main() {
    if ok, _ := regexp.MatchString("^-?[0-9a-zA-Z_]+$", "-hello123"); ok {
        fmt.Println("ok")
    } else {
        fmt.Println("no")
    }
}
