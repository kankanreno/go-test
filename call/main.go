package main

import (
	"fmt"
)

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}

type Intf interface {
	M1()
	M2()
}

func main() {
	t1 := T{"t1"}

	fmt.Println("M1调用前，t1.Name：", t1.Name)
	t1.M1()
	fmt.Println("M1调用后，t1.Name：", t1.Name)

	fmt.Println("M2调用前，t1.Name：", t1.Name)
	t1.M2()
	fmt.Println("M2调用后，t1.Name：", t1.Name)

	t2 := &T{"t2"}

	fmt.Println("M1调用前，t2.Name：", t2.Name)
	t2.M1()
	fmt.Println("M1调用后，t2.Name：", t2.Name)

	fmt.Println("M2调用前，t2.Name：", t2.Name)
	t2.M2()
	fmt.Println("M2调用后，t2.Name：", t2.Name)

	var t3 Intf = t2

	fmt.Println("M1调用前，t3.Name：", t3.Name)
	t3.M1()
	fmt.Println("M1调用后，t3.Name：", t3.Name)

	fmt.Println("M2调用前，t3.Name：", t3.Name)
	t3.M2()
	fmt.Println("M2调用后，t3.Name：", t3.Name)
}