package main

import "github.com/kecci/go-design-pattern/behavioral_patterns/observer/usecase"

func main() {

	shirtItem := usecase.NewItem("Nike Shirt")

	observerFirst := &usecase.Customer{Id: "abc@gmail.com"}
	observerSecond := &usecase.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
