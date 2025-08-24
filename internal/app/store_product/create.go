package store_product

import (
	"otel-prometheus-study/internal/domain/store_product"
	"otel-prometheus-study/internal/infra/postgres"
)

func CreateStoreProduct(storeID int, productID int, price string, quantity int) (store_product.StoreProduct, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return store_product.StoreProduct{}, err
	}
	defer db.Close()

	repo := postgres.NewStoreProductRepository(db)
	newStoreProduct, err := store_product.NewStoreProduct(storeID, productID, price, quantity)
	if err != nil {
		return store_product.StoreProduct{}, err
	}

	return repo.InsertStoreProduct(newStoreProduct)
}
