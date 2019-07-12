package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Println("m: ", numCPUs)
	tunny.NewFunc(numCPUs,)
}
