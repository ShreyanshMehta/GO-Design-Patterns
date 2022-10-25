package main

import "fmt"

func main() {
	var vehicleType string
	var vehicleCompany string
	fmt.Println("Enter your vehicle type and company")
	fmt.Scanln(&vehicleType, &vehicleCompany)
	factory, _ := createFactory(vehicleCompany)
	vehicle, _ := factory.createVehicle(vehicleType)
	vehicle.drive()
	vehicle.refill()
}
