package postgres

import (
	"database/sql"
	"otel-prometheus-study/internal/domain/product"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/logger"
	"strings"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{connection: db}
}

func (pr ProductRepository) InsertProduct(p product.Product) (product.Product, error) {
	query := `INSERT INTO Products (Name, Price) VALUES ($1, $2) RETURNING Id`
	name := strings.ToLower(p.Name())
	price := p.PriceValue.Value()

	var id int
	err := pr.connection.QueryRow(query, name, price).Scan(&id)
	if err != nil {
		logger.LogError(err, "context", "InsertProduct")
		return product.Product{}, err
	}

	newID, err := shared.NewID(id)
	if err != nil {
		return product.Product{}, err
	}
	p.IDValue = newID
	return p, nil
}

func (pr ProductRepository) GetProducts() ([]product.Product, error) {
	query := `SELECT Id, Name, Price FROM Products`
	rows, err := pr.connection.Query(query)
	if err != nil {
		logger.LogError(err, "context", "GetProducts")
		return nil, err
	}
	defer rows.Close()

	var list []product.Product
	for rows.Next() {
		var id int
		var name, price string
		if err := rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}

		prod, err := product.NewProduct(id, name, price)
		if err != nil {
			return nil, err
		}
		list = append(list, prod)
	}
	return list, nil
}
