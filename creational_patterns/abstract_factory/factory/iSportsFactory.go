package factory

import (
	"fmt"

	"github.com/kecci/design-pattern-go/creational_patterns/abstract_factory/factory/adidas"
	"github.com/kecci/design-pattern-go/creational_patterns/abstract_factory/factory/nike"
	"github.com/kecci/design-pattern-go/creational_patterns/abstract_factory/item"
)

type iSportsFactory interface {
	MakeShoe() item.IShoe
	MakeShirt() item.IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Adidas{}, nil
	}

	if brand == "nike" {
		return &nike.Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
