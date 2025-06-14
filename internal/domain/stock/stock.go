package stock

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

type Stock struct {
	id       shared.ID
	product  shared.ID
	quantity int
}

func NewStock(id, productID int, quantity int) (Stock, error) {
	if quantity < 0 {
		return Stock{}, fmt.Errorf("NewStock(): quantity must be non-negative")
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
		id:       stockID,
		product:  pID,
		quantity: quantity,
	}, nil
}

func (s Stock) ID() int {
	return s.id.Value()
}

func (s Stock) ProductID() int {
	return s.product.Value()
}

func (s Stock) Quantity() int {
	return s.quantity
}

func (s Stock) String() string {
	return fmt.Sprintf("{'id': '%d', 'product_id': '%d', 'quantity': %d}", s.ID(), s.ProductID(), s.Quantity())
}
