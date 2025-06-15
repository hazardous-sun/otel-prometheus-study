package product

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// Product
// Represents a product entity with ID, Name, and Price.
type Product struct {
	IDValue    shared.ID    `json:"id"`
	NameValue  shared.Name  `json:"name"`
	PriceValue shared.Price `json:"price"`
}

func (p Product) ID() int {
	return p.IDValue.Value()
}

func (p Product) Name() string {
	return p.NameValue.String()
}

func (p Product) Price() string {
	return p.PriceValue.String()
}

func (p Product) String() string {
	return fmt.Sprintf("{'ID': '%d', 'Name': '%s', 'Price': '%s'}", p.IDValue.Value(), p.NameValue.Value(), p.PriceValue.Value())
}

// NewProduct
// Creates a new Product instance, validating ID, Name, and Price inputs.
func NewProduct(inputID int, inputName, inputPrice string) (Product, error) {
	// Check if inputID is a valid ID
	id, err := shared.NewID(inputID)

	if err != nil {
		return Product{}, err
	}

	// Check if inputName is a valid Name
	name, err := shared.NewName(inputName)

	if err != nil {
		return Product{}, err
	}

	// Check if inputPrice is a valid Price
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
