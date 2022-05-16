package main

import (
	"fmt"
	"log"

	"github.com/kecci/go-design-pattern/structural_patterns/facade/usecase"
)

func main() {
	fmt.Println()
	walletFacade := usecase.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
