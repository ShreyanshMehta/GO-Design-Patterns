package main

import (
	"fmt"
	"sync"
)

type Vehicle struct{}

var VehicleObj *Vehicle
var VehicleLock = &sync.Mutex{}

func createVehicle() *Vehicle {
	if VehicleObj == nil {
		VehicleLock.Lock()
		defer VehicleLock.Unlock()
		if VehicleObj == nil {
			fmt.Printf("Creating a vehicle instance\n")
			VehicleObj = &Vehicle{}
		}
		return VehicleObj
	}
	return VehicleObj
}

func main() {
	for i := 0; i < 100; i++ {
		go createVehicle()
		fmt.Printf("Created Vehicle for %v times\n", i)
	}
}
