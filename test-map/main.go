package main

import "fmt"

/*func main() {
    m := make(map[int]int)

    go func() {
        for {
            _ = m[1]
        }
    }()

    go func() {
        for {
            m[2] = 2
        }
    }()

    select {}
}*/

/*func main() {
    var counter = struct {
        sync.RWMutex
        m map[string]int
    }{
        m: make(map[string]int),
    }

    go func() {
        for {
            counter.RLock()
            n := counter.m["some_key"]
            counter.RUnlock()
            fmt.Println("some_key", n)
        }
    }()

    go func() {
        for {
            counter.Lock()
            counter.m["some_key"]++
            counter.Unlock()
        }
    }()

    select {}
}*/

// https://studygolang.com/articles/10511

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
