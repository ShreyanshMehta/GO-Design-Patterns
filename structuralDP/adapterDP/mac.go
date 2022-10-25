package main

import "fmt"

type Mac struct {
	IComputer
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
