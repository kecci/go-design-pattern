package main

import (
	"fmt"

	"github.com/kecci/design-pattern-go/behavioral_patterns/iterator/usecase"
)

func main() {

	user1 := &usecase.User{
		Name: "a",
		Age:  30,
	}
	user2 := &usecase.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &usecase.UserCollection{
		Users: []*usecase.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
