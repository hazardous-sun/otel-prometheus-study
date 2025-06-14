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
	id            shared.ID
	name          shared.Name
	storeProducts []store_product.StoreProduct
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

	return Store{id: id, name: name, storeProducts: []store_product.StoreProduct{}}, nil
}

func (s Store) ID() int                                { return s.id.Value() }
func (s Store) Name() string                           { return s.name.Value() }
func (s Store) Products() []store_product.StoreProduct { return s.storeProducts }

// AddProduct
// Adds a StoreProduct to the store's inventory.
// Returns an error if the StoreProduct does not belong to this store.
func (s *Store) AddProduct(sp store_product.StoreProduct) error {
	if sp.StoreID() != s.ID() {
		return fmt.Errorf("AddProduct(): StoreProduct does not belong to this store")
	}
	s.storeProducts = append(s.storeProducts, sp)
	return nil
}

// String returns a formatted string representation of the Store.
func (s Store) String() string {
	return fmt.Sprintf("{'id': '%d', 'name': '%s', 'products_count': %d}", s.ID(), s.Name(), len(s.storeProducts))
}
