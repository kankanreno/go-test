package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer in main EXEC")

	fmt.Println("main EXEC")
	//os.Exit(1)
    panic("main EXEC")
	fooFunc()
}

func fooFunc() string {
	defer fmt.Println("defer in fooFunc EXEC")

	fmt.Println("fooFunc EXEC")
	//os.Exit(1)
    panic("fooFunc EXEC")
	return barFunc()
}

func barFunc() string {
	fmt.Println("barFunc EXEC")
	//os.Exit(1)
    panic("barFunc EXEC")
	return ""
}
