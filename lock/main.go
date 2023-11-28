package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SafeMap struct {
	m    map[string]string
	lock sync.Mutex
}

func (sm *SafeMap) Set(key, value string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) string {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	return sm.m[key]
}

func main() {
	sm := &SafeMap{m: make(map[string]string)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(strconv.Itoa(i), strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
	fmt.Println(sm.Get("99"))
}
