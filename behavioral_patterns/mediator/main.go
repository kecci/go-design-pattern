package main

import "github.com/kecci/design-pattern-go/behavioral_patterns/mediator/usecase"

func main() {
	stationManager := usecase.NewStationManger()

	passengerTrain := &usecase.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &usecase.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
