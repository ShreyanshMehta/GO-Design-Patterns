package main

import "fmt"

func main() {
	var comp1, comp2 IComputer
	var printer1, printer2 Printer
	printer1 = &Hp{}
	printer2 = &Epson{}
	comp1 = &Mac{}
	comp2 = &Windows{}

	comp1.SetPrinter(printer1)
	comp1.Print()
	fmt.Println()

	comp1.SetPrinter(printer2)
	comp1.Print()
	fmt.Println()

	comp2.SetPrinter(printer1)
	comp2.Print()
	fmt.Println()

	comp2.SetPrinter(printer2)
	comp2.Print()
	fmt.Println()
}
