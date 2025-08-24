package store

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/domain/store_product"
)

// Store
// Represents a physical or virtual store entity.
// Each store has an ID, a Name, and a list of StoreProduct items.
type Store struct {
	IDValue       shared.ID                    `json:"id"`
	NameValue     shared.Name                  `json:"name"`
	StoreProducts []store_product.StoreProduct `json:"storeProductsIDs"`
}

// NewStore
// Constructs a new Store with validated ID and Name.
// Returns an error if validation fails.
func NewStore(inputID int, inputName string) (Store, error) {
	id, err := shared.NewID(inputID)
	if err != nil {
		return Store{}, err
	}

	name, err := shared.NewName(inputName)
	if err != nil {
		return Store{}, err
	}

	return Store{IDValue: id, NameValue: name, StoreProducts: []store_product.StoreProduct{}}, nil
}

func (s Store) ID() int                                { return s.IDValue.Value() }
func (s Store) Name() string                           { return s.NameValue.Value() }
func (s Store) Products() []store_product.StoreProduct { return s.StoreProducts }

// AddProduct
// Adds a StoreProduct to the store's inventory.
// Returns an error if the StoreProduct does not belong to this store.
func (s Store) AddProduct(sp store_product.StoreProduct) (Store, error) {
	if sp.StoreID() != s.ID() {
		return s, fmt.Errorf("AddProduct(): StoreProduct does not belong to this store")
	}
	s.StoreProducts = append(s.StoreProducts, sp)
	return s, nil
}

// String returns a formatted string representation of the Store.
func (s Store) String() string {
	return fmt.Sprintf("{'ID': '%d', 'Name': '%s', 'products_count': %d}", s.ID(), s.Name(), len(s.StoreProducts))
}
