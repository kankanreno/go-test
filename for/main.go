package main

import "fmt"

func main() {
	stus := []struct {
		Name string
		Age  int
	}{
		{"宋", 22},
		{"高", 23},
		{"徐", 24},
		{"李", 25},
	}

	for i := range stus {
		stus[i].Name = "x"
	}

	fmt.Print(stus)

}
