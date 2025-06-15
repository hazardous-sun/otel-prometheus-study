package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/domain/stock"
	"otel-prometheus-study/internal/logger"
)

type StockRepository struct {
	connection *sql.DB
}

func (sr StockRepository) InsertStock(newStock stock.Stock) (stock.Stock, error) {
	stockId := newStock.IDValue.Value()
	productId := newStock.ProductIDValue.Value()
	quantity := newStock.Quantity()
	logger.LogDebug("Preparing insert query", "id", stockId, "product_id", productId, "quantity", quantity)

	query, err := sr.connection.Prepare(`INSERT INTO stock (id, store_id, product_id, quantity) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		logger.LogError(err, "context", "preparing insert statement")
		if errors.Is(err, sql.ErrNoRows) {
			return stock.Stock{}, errors.New("stocks table missing")
		}
		return stock.Stock{}, err
	}
	defer query.Close()

	var id int

	err = query.QueryRow(stockId, stockId, productId, quantity).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			logger.LogWarning("Postgres error", "code", pqErr.Code, "constraint", pqErr.Constraint)
			if pqErr.Code == "23505" && pqErr.Constraint == "customers_name_key" {
				return stock.Stock{}, errors.New("customer already exists")
			}
		}
		logger.LogError(err, "context", "executing insert query")
		return stock.Stock{}, err
	}

	newID, err := shared.NewID(id)
	if err != nil {
		logger.LogError(err, "context", "constructing shared.ID")
		return stock.Stock{}, err
	}

	newStock.IDValue = newID
	logger.LogSuccess("Stock inserted successfully", "id", newID.Value(), "product_id", newStock.ProductIDValue, "quantity", newStock.Quantity())
	return newStock, nil
}

func (sr StockRepository) GetStockByID(id int) (stock.Stock, error) {
	logger.LogDebug("Preparing select query", "id", id)

	query, err := sr.connection.Prepare("SELECT id, name FROM customers WHERE id = $1")
	if err != nil {
		logger.LogError(err, "context", "preparing select statement")
		if errors.Is(err, sql.ErrNoRows) {
			return stock.Stock{}, errors.New("stock not found")
		}
		return stock.Stock{}, err
	}
	defer query.Close()

	var productID, quantity int

	if err = query.QueryRow(id).Scan(&id, &productID, &quantity); err != nil {
		logger.LogError(err, "context", "executing select query")
		return stock.Stock{}, err
	}

	stockObj, err := stock.NewStock(id, productID, quantity)
	if err != nil {
		logger.LogError(err, "context", "constructing stock")
		return stock.Stock{}, err
	}

	logger.LogSuccess("Stock retrieved successfully", "id", stockObj.ID(), "product_id", stockObj.ProductID(), "quantity", stockObj.Quantity())
	return stockObj, nil
}

func (sr StockRepository) GetStockByProductID(productID int) (stock.Stock, error) {
	logger.LogDebug("Preparing select query", "id", productID)

	query, err := sr.connection.Prepare("SELECT id, name FROM customers WHERE id = $1")
	if err != nil {
		logger.LogError(err, "context", "preparing select statement")
		if errors.Is(err, sql.ErrNoRows) {
			return stock.Stock{}, errors.New("stock not found")
		}
		return stock.Stock{}, err
	}
	defer query.Close()

	var id, quantity int

	if err = query.QueryRow(id).Scan(&id, &productID, &quantity); err != nil {
		logger.LogError(err, "context", "executing select query")
		return stock.Stock{}, err
	}

	stockObj, err := stock.NewStock(id, productID, quantity)
	if err != nil {
		logger.LogError(err, "context", "constructing stock")
		return stock.Stock{}, err
	}

	logger.LogSuccess("Stock retrieved successfully", "id", stockObj.ID(), "product_id", stockObj.ProductID(), "quantity", stockObj.Quantity())
	return stockObj, nil
}

func (sr StockRepository) GetStocks() ([]stock.Stock, error) {
	query := "SELECT id, product_id, quantity FROM stocks"

	rows, err := sr.connection.Query(query)
	if err != nil {
		logger.LogError(err, "context", "executing select query")
		return []stock.Stock{}, err
	}
	defer rows.Close()

	var stockList []stock.Stock
	var stockObj stock.Stock

	for rows.Next() {
		err = rows.Scan(
			&stockObj.IDValue,
			&stockObj.ProductIDValue,
			&stockObj.QuantityValue,
		)

		if err != nil {
			logger.LogError(err, "context", "scanning stock row")
			return []stock.Stock{}, err
		}

		stockList = append(stockList, stockObj)
	}

	if err = rows.Err(); err != nil {
		logger.LogError(err, "context", "iterating over stock rows")
		return []stock.Stock{}, err
	}

	return stockList, nil
}

func NewStockRepository(db *sql.DB) StockRepository {
	return StockRepository{connection: db}
}
