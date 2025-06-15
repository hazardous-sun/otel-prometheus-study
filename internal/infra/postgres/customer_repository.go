package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"otel-prometheus-study/internal/domain/customer"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/logger" // import your custom logger
	"strings"
)

type CustomerRepository struct {
	connection *sql.DB
}

func (cr CustomerRepository) InsertCustomer(customer customer.Customer) (customer.Customer, error) {
	name := strings.ToLower(customer.Name())
	logger.LogDebug("Preparing insert query", "name", name)

	query, err := cr.connection.Prepare("INSERT INTO customers (name) VALUES ($1) RETURNING id")
	if err != nil {
		logger.LogError(err, "context", "preparing insert statement")
		if errors.Is(err, sql.ErrNoRows) {
			return customer, errors.New("customers table missing")
		}
		return customer, err
	}
	defer query.Close()

	var id int
	err = query.QueryRow(name).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			logger.LogWarning("Postgres error", "code", pqErr.Code, "constraint", pqErr.Constraint)
			if pqErr.Code == "23505" && pqErr.Constraint == "customers_name_key" {
				return customer, errors.New("customer already exists")
			}
		}
		logger.LogError(err, "context", "executing insert query")
		return customer, err
	}

	newID, err := shared.NewID(id)
	if err != nil {
		logger.LogError(err, "context", "constructing shared.ID")
		return customer, err
	}

	customer.IDValue = newID
	logger.LogSuccess("Customer inserted successfully", "id", newID.Value(), "name", name)
	return customer, nil
}
