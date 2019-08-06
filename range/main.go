package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("============== 普通 slice")
	INT_LIST := []int{1, 2, 3}
	fmt.Printf("&INT_LIST[0]: %p\n", &INT_LIST[0])

	fmt.Println("====")
	for _, v := range INT_LIST {
		fmt.Printf("v type: %T\n", v)
		fmt.Printf("&v: %p\n", &v)
	}

	fmt.Println("====")
	for _, v := range INT_LIST {
		_v := v
		fmt.Printf("_v type: %T\n", _v)
		fmt.Printf("&_v: %p\n", &_v)
	}

	fmt.Println("====")
	for i := range INT_LIST {
		fmt.Printf("INT_LIST[%d] type: %T\n", i, INT_LIST[i])
		fmt.Printf("&INT_LIST[%d]: %p\n", i, &INT_LIST[i])
	}

	fmt.Println("============== 结构体 slice")
	stus := []struct {
		ID   int
		Name string
	}{
		{
			1,
			"foo",
		}, {
			2,
			"bar",
		}, {
			3,
			"baz",
		},
	}
	fmt.Printf("%p\n", &stus[0])

	fmt.Println("====")
	for _, v := range stus {
		fmt.Printf("v type: %T\n", v)
		fmt.Printf("%p\n", &v)
		v.ID = 100
		v.Name = "qux"
	}
	fmt.Printf("stus: %+v\n", stus)

	fmt.Println("============== map slice")
	SALES_CHANNEL_LIST := []map[string]interface{}{
		{
			"ID":   1,
			"Name": "foo",
		}, {
			"ID":   2,
			"Name": "bar",
		}, {
			"ID":   3,
			"Name": "baz",
		},
	}
	fmt.Printf("%p\n", SALES_CHANNEL_LIST[0])

	fmt.Println("====")
	for _, v := range SALES_CHANNEL_LIST {
		fmt.Printf("v type: %T\n", v)
		fmt.Printf("&v: %p\n", &v)
		v["ID"] = 100
		v["Name"] = "qux"
	}
	fmt.Println(SALES_CHANNEL_LIST)


	fmt.Println("============== range & goroutine")
	wg := sync.WaitGroup{}
	//intList := []int{1, 2, 3, 4, 5}
	//for i := range intList {
	//
	//}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}