package main

import "C"
import "fmt"

type Car struct {
	Brand	string
	Type	string
	Color 	string
}

func (car *Car) Drive() {
	fmt.Printf("A %s %s %s car is running on the road!\n", car.Color, car.Type, car.Brand)
}

type Builder interface {
	AddBrand(string) Builder
	SetType(string) Builder
	PaintColor(string) Builder
	GetResult() Car
}

type ConcreteBuilder struct {
	C Car
}

func (concreteBuilder *ConcreteBuilder) AddBrand(brand string) Builder {
	concreteBuilder.C.Brand = brand
	return concreteBuilder
}

func (concreteBuilder *ConcreteBuilder) SetType(tp string) Builder {
	concreteBuilder.C.Type = tp
	return concreteBuilder
}

func (concreteBuilder *ConcreteBuilder) PaintColor(color string) Builder {
	concreteBuilder.C.Color = color
	return concreteBuilder
}

func (concreteBuilder *ConcreteBuilder) GetResult() Car {
	return concreteBuilder.C
}

type Director struct {
	B Builder
}

func main() {
	director := Director{&ConcreteBuilder{}}
	adCar := director.B.SetType("SUV").AddBrand("奥迪").PaintColor("white").GetResult()
	adCar.Drive()

	bwCar := director.B.SetType("sporting").AddBrand("宝马").PaintColor("red").GetResult()
	bwCar.Drive()
}















