package product

import (
	"otel-prometheus-study/internal/domain/product"
	"otel-prometheus-study/internal/infra/postgres"
)

func CreateProduct(name string, price string) (product.Product, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return product.Product{}, err
	}
	defer db.Close()

	repo := postgres.NewProductRepository(db)
	newProduct, err := product.NewProduct(0, name, price)
	if err != nil {
		return product.Product{}, err
	}

	return repo.InsertProduct(newProduct)
}
