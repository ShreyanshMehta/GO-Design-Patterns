package main

import "fmt"

type Car struct {
	IVehicle
}

func newCar() IVehicle {
	return &Car{}
}

func (c *Car) drive() {
	fmt.Println("driving a car")
}

func (c *Car) refill(quantity int) {
	fmt.Printf("refilling a car: %v\n", quantity)
}
