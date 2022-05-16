package main

import (
	"fmt"

	"github.com/kecci/go-design-pattern/creational_patterns/singleton/usecase"
)

func main() {

	for i := 0; i < 30; i++ {
		go usecase.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
