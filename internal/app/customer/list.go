package customer

import (
	"otel-prometheus-study/internal/domain/customer"
	"otel-prometheus-study/internal/infra/postgres"
)

func ListCustomers() ([]customer.Customer, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := postgres.NewCustomerRepository(db)
	return repo.GetCustomers()
}
