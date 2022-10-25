package main

import "fmt"

type IVehicle interface {
	drive()
	refill()
}

type TataCar struct {
	IVehicle
}

func (c *TataCar) drive() {
	fmt.Println("Driving Tata car")
}

func (c *TataCar) refill() {
	fmt.Println("Refilling Tata car")
}

type JeepCar struct {
	IVehicle
}

func (c *JeepCar) drive() {
	fmt.Println("Driving Jeep car")
}

func (c *JeepCar) refill() {
	fmt.Println("Refilling Jeep car")
}

type TataBike struct {
	IVehicle
}

func (c *TataBike) drive() {
	fmt.Println("Driving Tata Bike")
}

func (c *TataBike) refill() {
	fmt.Println("Refilling Tata Bike")
}

type JeepBike struct {
	IVehicle
}

func (c *JeepBike) drive() {
	fmt.Println("Driving Jeep Bike")
}

func (c *JeepBike) refill() {
	fmt.Println("Refilling Jeep Bike")
}
