package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/structural_patterns/bridge/usecase"
)

func main() {

	hpPrinter := &usecase.Hp{}
	epsonPrinter := &usecase.Epson{}

	macComputer := &usecase.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &usecase.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
