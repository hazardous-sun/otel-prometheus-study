package product

import "otel-prometheus-study/internal/domain/shared"

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
