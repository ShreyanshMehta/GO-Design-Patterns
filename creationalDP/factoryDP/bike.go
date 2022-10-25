package main

import "fmt"

type Bike struct {
	IVehicle
}

func newBike() IVehicle {
	return &Bike{}
}

func (c *Bike) drive() {
	fmt.Println("driving a bike")
}

func (c *Bike) refill(quantity int) {
	fmt.Printf("refilling a bike: %v\n", quantity)
}
