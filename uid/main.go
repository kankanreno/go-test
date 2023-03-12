package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {

	for i := 0; i < 1000000; i++ {
		id := xid.New()
		fmt.Println(id)
	}

}
