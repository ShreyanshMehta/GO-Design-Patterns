package main

import "fmt"

func main() {
	var vehicleType string
	fmt.Println("Enter the vehicle type")
	_, err := fmt.Scanln(&vehicleType)
	if err != nil {
		return
	}
	v, _ := createVehicle(vehicleType)
	v.refill(5)
	v.drive()
}
