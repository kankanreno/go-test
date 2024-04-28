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
	m := map[string][]string{"name": {"看看"}, "mail": {"kankan@pa.com"}, "phone": {"18352515222"}}
	m2 := map[string]string{"name": "看看", "mail": "kankan@pa.com", "phone": "18352515222"}

	// m
	fmt.Println("m: ", m)
	fmt.Println("name: ", m["name"][0])
	//fmt.Println("name: ", m["name2"][0])

	// m2
	fmt.Println("m: ", m2)
	fmt.Println("name: ", m2["name"])
	fmt.Println("name2: ", m2["name2"])

	m2["name"] = ""
	fmt.Printf("m2: %+v\n", m2)

	delete(m2, "name")
	fmt.Printf("m2: %+v\n", m2)
}
