package main

import (
	"errors"
	"fmt"
)

func createVehicle(vehicleType string) (IVehicle, error) {
	var vehicle IVehicle
	fmt.Printf("Creating vehicle: %v\n", vehicleType)
	if vehicleType == "car" {
		vehicle = newCar()
	} else if vehicleType == "bike" {
		vehicle = newBike()
	} else {
		return nil, errors.New("wrong vehicle type\n")
	}
	return vehicle, nil
}
