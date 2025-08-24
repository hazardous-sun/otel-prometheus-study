package stock

import (
	"otel-prometheus-study/internal/domain/stock"
	"otel-prometheus-study/internal/infra/postgres"
)

func ListStocks() ([]stock.Stock, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := postgres.NewStockRepository(db)
	return repo.GetStocks()
}
