package main

import "fmt"

type Windows struct {
	IComputer
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
