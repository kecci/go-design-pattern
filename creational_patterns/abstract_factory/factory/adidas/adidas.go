package adidas

import "github.com/kecci/go-design-pattern/creational_patterns/abstract_factory/item"

type Adidas struct {
}

func (a *Adidas) MakeShoe() item.IShoe {
	return &adidasShoe{
		Shoe: item.Shoe{
			Logo: "adidas",
			Size: 44,
		},
	}
}

func (a *Adidas) MakeShirt() item.IShirt {
	return &adidasShirt{
		Shirt: item.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
