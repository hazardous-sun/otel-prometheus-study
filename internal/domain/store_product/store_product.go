package store_product

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// StoreProduct
// Represents the association of a product with a store,
// including the store ID, product ID, specific price, and quantity available.
type StoreProduct struct {
	storeID   shared.ID
	productID shared.ID
	price     shared.Price
	quantity  int
}

// NewStoreProduct
// Creates a new StoreProduct after validating the storeID, productID, price, and quantity.
// Returns an error if any validation fails (e.g., negative quantity or invalid price).
func NewStoreProduct(storeID, productID int, price string, quantity int) (StoreProduct, error) {
	if quantity < 0 {
		return StoreProduct{}, fmt.Errorf("NewStoreProduct(): quantity must be non-negative")
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
		storeID:   sID,
		productID: pID,
		price:     pr,
		quantity:  quantity,
	}, nil
}

func (sp StoreProduct) StoreID() int   { return sp.storeID.Value() }
func (sp StoreProduct) ProductID() int { return sp.productID.Value() }
func (sp StoreProduct) Price() string  { return sp.price.String() }
func (sp StoreProduct) Quantity() int  { return sp.quantity }

func (sp StoreProduct) String() string {
	return fmt.Sprintf(
		"{'store_id': '%d', 'product_id': '%d', 'price': '%s', 'quantity': %d}",
		sp.StoreID(), sp.ProductID(), sp.Price(), sp.Quantity(),
	)
}
