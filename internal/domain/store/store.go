package store

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/domain/store_product"
)

type Store struct {
	id            shared.ID
	name          shared.Name
	storeProducts []store_product.StoreProduct
}

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

func (s *Store) AddProduct(sp store_product.StoreProduct) error {
	if sp.StoreID() != s.ID() {
		return fmt.Errorf("AddProduct(): StoreProduct does not belong to this store")
	}
	s.storeProducts = append(s.storeProducts, sp)
	return nil
}

func (s Store) String() string {
	return fmt.Sprintf("{'id': '%d', 'name': '%s', 'products_count': %d}", s.ID(), s.Name(), len(s.storeProducts))
}
