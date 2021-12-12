package main

import "fmt"

type Bmw struct {
	model    string
	price    int
	maxSpeed int
}

func NewBmw(model string, price int, maxSpeed int) *Bmw {
	return &Bmw{model: model, price: price, maxSpeed: maxSpeed}
}

type BmwFactory struct {
}

func (bwF BmwFactory) create(model string) *Bmw {
	switch model {
	case "X5":
		return NewBmw(model, 700300, 180)
	case "X6":
		return NewBmw(model, 800000, 200)
	default:
		return NewBmw(model, 0, 0)
	}
}
func main() {
	factory := new(BmwFactory)
	fmt.Println(factory.create("X5"))
}
