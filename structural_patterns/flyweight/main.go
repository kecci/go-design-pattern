package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/structural_patterns/flyweight/usecase"
)

func main() {
	game := usecase.NewGame()

	//Add Terrorist
	game.AddTerrorist(usecase.TerroristDressType)
	game.AddTerrorist(usecase.TerroristDressType)
	game.AddTerrorist(usecase.TerroristDressType)
	game.AddTerrorist(usecase.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(usecase.CounterTerrroristDressType)
	game.AddCounterTerrorist(usecase.CounterTerrroristDressType)
	game.AddCounterTerrorist(usecase.CounterTerrroristDressType)

	dressFactoryInstance := usecase.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
