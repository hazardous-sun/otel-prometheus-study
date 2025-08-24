package stock

import (
	"otel-prometheus-study/internal/domain/stock"
	"otel-prometheus-study/internal/infra/postgres"
)

func CreateStock(productID int, quantity int) (stock.Stock, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return stock.Stock{}, err
	}
	defer db.Close()

	repo := postgres.NewStockRepository(db)
	newStock, err := stock.NewStock(0, productID, quantity)
	if err != nil {
		return stock.Stock{}, err
	}

	return repo.InsertStock(newStock)
}
