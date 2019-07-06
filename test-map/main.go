package main

import "fmt"

func main() {
	m := map[string][]string{"cn": {"看看"}, "mail": {"kankan@pa.com"}, "mobile": {"18352515222"}}
	m2 := map[string]string{"cn": "看看", "mail": "kankan@pa.com", "mobile": "18352515222"}

	// m
	fmt.Println("m: ", m)
	fmt.Println("cn: ", m["cn"][0])
	//fmt.Println("cn: ", m["cn2"][0])

	// m2
	fmt.Println("m: ", m2)
	fmt.Println("cn: ", m2["cn"])
	fmt.Println("cn: ", m2["cn2"])
}
