package product

import (
	"otel-prometheus-study/internal/domain/product"
	"otel-prometheus-study/internal/infra/postgres"
)

func ListProducts() ([]product.Product, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := postgres.NewProductRepository(db)
	return repo.GetProducts()
}
