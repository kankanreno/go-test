package main

import "fmt"

func main() {
	//var w Wheel
	//w.X = 8
	//w.Point.X = 9

	w := &Wheel{
		Circle{
			Point{
				X: 88,
				Y: 99,
			},
		},
	}


	w.print()
	//fmt.Printf("s1 = %+v\n", w)
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
}

type Wheel struct {
	Circle
}

func (p Point) print() {
	fmt.Println("PPrint: ", p.X, p.Y)
}


func (c Circle) print() {
	fmt.Println("CPrint: ", c.Point.X, c.Point.Y)
}

