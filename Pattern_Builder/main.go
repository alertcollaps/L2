package main

import "fmt"

type Car struct {
	doors int
	nitro bool
	wheels int
	length int
	weight int
	spoiler bool
}

func NewCar() *Car{
	cr := new(Car)
	cr.doors = 4
	cr.nitro = false
	cr.wheels = 4
	cr.spoiler = true
	return cr
}
func (c Car)PrintInfo()  {
	fmt.Printf("doors = %v\nnitro: %v\n" +
		"wheels = %v\nspoiler:%v", c.doors, c.nitro, c.wheels, c.spoiler)
}

type CarBuilder struct {
	car *Car
}
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{NewCar()}
}

func (CrBd *CarBuilder)setNitro(nitro bool) *CarBuilder{
	CrBd.car.nitro = nitro
	return CrBd
}

func (CrBd *CarBuilder)setDoors(doors int) *CarBuilder{
	CrBd.car.doors = doors
	return CrBd
}

func (CrBd *CarBuilder)setWheels(wheels int) *CarBuilder{
	CrBd.car.wheels = wheels
	return CrBd
}

func (CrBd *CarBuilder)setSpoiler(spoiler bool) *CarBuilder{
	CrBd.car.spoiler = spoiler
	return CrBd
}

func (CrBd *CarBuilder)Build() *CarBuilder{
	return CrBd
}

type Director struct {

}

func (dir *Director)Passenger_Car() *CarBuilder{
	CrBd := NewCarBuilder().setSpoiler(false).Build()
	return CrBd
}

func (dir *Director)Racing_Car() *CarBuilder{
	CrBd := NewCarBuilder().setNitro(true).setDoors(2).Build()
	return CrBd
}

func (dir *Director)Freight_Car() *CarBuilder{
	CrBd := NewCarBuilder().setSpoiler(false).
		setDoors(2).
		setWheels(8).
		Build()
	return CrBd
}

func main() {
	dr := new(Director)
	Freightcar := dr.Freight_Car()
	Freightcar.car.PrintInfo()
}
