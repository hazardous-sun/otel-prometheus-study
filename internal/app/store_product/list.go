package store_product

import (
	"otel-prometheus-study/internal/domain/store_product"
	"otel-prometheus-study/internal/infra/postgres"
)

func ListStoreProducts() ([]store_product.StoreProduct, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := postgres.NewStoreProductRepository(db)
	return repo.GetStoreProducts()
}
