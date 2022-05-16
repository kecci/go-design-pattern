package main

import "github.com/kecci/go-design-pattern/behavioral_patterns/mediator/usecase"

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
