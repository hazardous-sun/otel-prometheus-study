package product

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

type Product struct {
	id    shared.ID
	name  shared.Name
	price shared.Price
}

func (p Product) Name() string {
	return p.name.String()
}

func (p Product) Price() string {
	return p.price.String()
}

func (p Product) String() string {
	return fmt.Sprintf("{'id': '%d', 'name': '%s', 'price': '%s'}", p.id.Value(), p.name.Value(), p.price.Value())
}

func NewProduct(inputID int, inputName, inputPrice string) (Product, error) {
	// Check if inputID is a valid ID
	id, err := shared.NewID(inputID)

	if err != nil {
		return Product{}, err
	}

	// Check if inputName is a valid name
	name, err := shared.NewName(inputName)

	if err != nil {
		return Product{}, err
	}

	// Check if inputPrice is a valid price
	price, err := shared.NewPrice(inputPrice)

	if err != nil {
		return Product{}, err
	}

	return Product{
		id,
		name,
		price,
	}, nil
}
