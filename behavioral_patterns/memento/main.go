package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/behavioral_patterns/memento/usecase"
)

func main() {

	caretaker := &usecase.Caretaker{
		MementoArray: make([]*usecase.Memento, 0),
	}

	originator := &usecase.Originator{
		State: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.Restorememento(caretaker.GetMenento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.Restorememento(caretaker.GetMenento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

}
