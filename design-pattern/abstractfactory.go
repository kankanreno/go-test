package main

import "fmt"

type AbstractProductA interface {
	Use()
}

type AbstractProductB interface {
	Eat()
}

type ProductA1 struct {

}

type ProductA2 struct {

}

type ProductB1 struct {

}

type ProductB2 struct {

}

func (*ProductA1) Use() {
	fmt.Println("Use Product A1")
}

func (*ProductA2) Use() {
	fmt.Println("Use Product A2")
}

func (*ProductB1) Eat() {
	fmt.Println("Eat Product B1")
}

func (*ProductB2) Eat() {
	fmt.Println("Eat Product B2")
}

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type ConcreteFactory1 struct {

}

type ConcreteFactory2 struct {

}

func (*ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ProductA1{}
}

func (*ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ProductB1{}
}

func (*ConcreteFactory2) CreateProductA() AbstractProductA {
	return &ProductA2{}
}

func (*ConcreteFactory2) CreateProductB() AbstractProductB {
	return &ProductB2{}
}

func main() {
	factory1 := &ConcreteFactory1{}
	factory2 := &ConcreteFactory2{}
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	productA1.Use()
	productB1.Eat()
	productA2.Use()
	productB2.Eat()
}























