package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/creational_patterns/factory_method/usecase"
)

func main() {
	ak47, _ := usecase.GetGun("ak47")
	musket, _ := usecase.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g usecase.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
