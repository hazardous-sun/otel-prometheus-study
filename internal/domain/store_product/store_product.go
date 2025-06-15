package store_product

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// StoreProduct
// Represents the association of a product with a store,
// including the store ID, product ID, specific Price, and Quantity available.
type StoreProduct struct {
	StoreIDValue   shared.ID    `json:"storeID"`
	ProductIDValue shared.ID    `json:"productID"`
	PriceValue     shared.Price `json:"price"`
	QuantityValue  int          `json:"quantity"`
}

// NewStoreProduct
// Creates a new StoreProduct after validating the StoreID, ProductID, Price, and Quantity.
// Returns an error if any validation fails (e.g., negative Quantity or invalid Price).
func NewStoreProduct(storeID, productID int, price string, quantity int) (StoreProduct, error) {
	if quantity < 0 {
		return StoreProduct{}, fmt.Errorf("NewStoreProduct(): Quantity must be non-negative")
	}

	sID, err := shared.NewID(storeID)
	if err != nil {
		return StoreProduct{}, err
	}

	pID, err := shared.NewID(productID)
	if err != nil {
		return StoreProduct{}, err
	}

	pr, err := shared.NewPrice(price)
	if err != nil {
		return StoreProduct{}, err
	}

	return StoreProduct{
		StoreIDValue:   sID,
		ProductIDValue: pID,
		PriceValue:     pr,
		QuantityValue:  quantity,
	}, nil
}

func (sp StoreProduct) StoreID() int   { return sp.StoreIDValue.Value() }
func (sp StoreProduct) ProductID() int { return sp.ProductIDValue.Value() }
func (sp StoreProduct) Price() string  { return sp.PriceValue.String() }
func (sp StoreProduct) Quantity() int  { return sp.QuantityValue }

func (sp StoreProduct) String() string {
	return fmt.Sprintf(
		"{'store_id': '%d', 'product_id': '%d', 'Price': '%s', 'Quantity': %d}",
		sp.StoreID(), sp.ProductID(), sp.Price(), sp.Quantity(),
	)
}
