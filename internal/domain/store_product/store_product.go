package store

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

type StoreProduct struct {
	storeID   shared.ID
	productID shared.ID
	price     shared.Price
	quantity  int
}

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
