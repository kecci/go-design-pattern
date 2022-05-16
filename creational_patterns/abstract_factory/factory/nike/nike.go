package nike

import "github.com/kecci/go-design-pattern/creational_patterns/abstract_factory/item"

type Nike struct {
}

func (n *Nike) MakeShoe() item.IShoe {
	return &nikeShoe{
		Shoe: item.Shoe{
			Logo: "nike",
			Size: 44,
		},
	}
}

func (n *Nike) MakeShirt() item.IShirt {
	return &nikeShirt{
		Shirt: item.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
