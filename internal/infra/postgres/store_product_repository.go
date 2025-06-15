package postgres

import (
	"database/sql"
	"otel-prometheus-study/internal/domain/store_product"
	"otel-prometheus-study/internal/logger"
)

type StoreProductRepository struct {
	connection *sql.DB
}

func (spr StoreProductRepository) InsertStoreProduct(sp store_product.StoreProduct) (store_product.StoreProduct, error) {
	query := `INSERT INTO store_products (store_id, product_id, price, quantity) VALUES ($1, $2, $3, $4)`
	_, err := spr.connection.Exec(query, sp.StoreID(), sp.ProductID(), sp.Price(), sp.Quantity())
	if err != nil {
		logger.LogError(err, "context", "InsertStoreProduct")
		return store_product.StoreProduct{}, err
	}
	return sp, nil
}

func (spr StoreProductRepository) GetStoreProducts() ([]store_product.StoreProduct, error) {
	query := `SELECT store_id, product_id, price, quantity FROM store_products`
	rows, err := spr.connection.Query(query)
	if err != nil {
		logger.LogError(err, "context", "GetStoreProducts")
		return nil, err
	}
	defer rows.Close()

	var list []store_product.StoreProduct
	for rows.Next() {
		var storeID, productID, quantity int
		var price string
		if err := rows.Scan(&storeID, &productID, &price, &quantity); err != nil {
			return nil, err
		}

		sp, err := store_product.NewStoreProduct(storeID, productID, price, quantity)
		if err != nil {
			return nil, err
		}
		list = append(list, sp)
	}
	return list, nil
}

func NewStoreProductRepository(db *sql.DB) StoreProductRepository {
	return StoreProductRepository{connection: db}
}
