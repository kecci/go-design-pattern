package main

import (
	"fmt"

	"github.com/kecci/design-pattern-go/behavioral_patterns/visitor/usecase"
)

func main() {
	square := &usecase.Square{Side: 2}
	circle := &usecase.Circle{Radius: 3}
	rectangle := &usecase.Rectangle{L: 2, B: 3}

	areaCalculator := &usecase.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &usecase.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
