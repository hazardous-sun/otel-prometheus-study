package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"otel-prometheus-study/internal/domain/customer"
	"otel-prometheus-study/internal/domain/shared"
	"strings"
)

type CustomerRepository struct {
	connection *sql.DB
}

func (cr CustomerRepository) InsertCustomer(customer customer.Customer) (customer.Customer, error) {
	query, err := cr.connection.Prepare("INSERT INTO customers (name) VALUES ($1) RETURNING id")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customer, errors.New("customers table missing")
		}
		return customer, err
	}

	var id int
	name := strings.ToLower(customer.Name())
	err = query.QueryRow(name).Scan(&id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" && pqErr.Constraint == "customers_name_key" {
				return customer, errors.New("customer already exists")
			}
		}
		return customer, err
	}

	err = query.Close()
	if err != nil {
		return customer, err
	}

	newID, err := shared.NewID(id)
	if err != nil {
		return customer, err
	}

	customer.IDValue = newID
	return customer, nil
}
