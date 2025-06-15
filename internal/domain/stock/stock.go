package stock

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// Stock
// Represents inventory for a Product.
// It includes a unique stock ID, a reference to the Product ID, and a Quantity.
type Stock struct {
	IDValue        shared.ID `json:"id"`
	ProductIDValue shared.ID `json:"product_id"`
	QuantityValue  int       `json:"quantity"`
}

// NewStock
// Creates a new Stock instance after validating the input.
func NewStock(id, productID int, quantity int) (Stock, error) {
	if quantity < 0 {
		return Stock{}, fmt.Errorf("NewStock(): Quantity must be non-negative")
	}

	stockID, err := shared.NewID(id)
	if err != nil {
		return Stock{}, err
	}

	pID, err := shared.NewID(productID)
	if err != nil {
		return Stock{}, err
	}

	return Stock{
		IDValue:        stockID,
		ProductIDValue: pID,
		QuantityValue:  quantity,
	}, nil
}

func (s Stock) ID() int {
	return s.IDValue.Value()
}

func (s Stock) ProductID() int {
	return s.ProductIDValue.Value()
}

func (s Stock) Quantity() int {
	return s.QuantityValue
}

func (s Stock) String() string {
	return fmt.Sprintf("{'ID': '%d', 'product_id': '%d', 'Quantity': %d}", s.ID(), s.ProductID(), s.Quantity())
}
