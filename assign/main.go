package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func main() {
	fmt.Println("============== wos")
	wos := []*struct {
		ID  uint
		UID uint
	}{
		{
			11,
			0,
		}, {
			12,
			0,
		}, {
			13,
			0,
		}, {
			14,
			0,
		},
	}
	fmt.Printf("%v\n", wos)

	fmt.Println("============== wos")
	users := []*struct {
		ID  uint
		Name string
	}{
		{
			501,
			"kankan",
		}, {
			502,
			"lili",
		}, {
			503,
			"binbin",
		},
	}
	fmt.Printf("%v\n", users)

	fmt.Println("============== assign")

	usersLen := len(users)
	for _, v := range wos {
		randInt := randomInt(usersLen)
		//fmt.Printf("%v\n", randInt)
		v.UID = users[randInt].ID
	}

	for _, v := range wos {
		fmt.Printf("%v\n", v)
	}
}
