package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/structural_patterns/decorator/usecase"
)

func main() {

	pizza := &usecase.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &usecase.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &usecase.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
