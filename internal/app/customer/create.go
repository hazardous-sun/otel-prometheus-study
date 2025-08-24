package customer

import (
	"otel-prometheus-study/internal/domain/customer"
	"otel-prometheus-study/internal/infra/postgres"
)

func CreateCustomer(name string) (customer.Customer, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return customer.Customer{}, err
	}
	defer db.Close()

	repo := postgres.NewCustomerRepository(db)
	newCustomer, err := customer.NewCustomer(0, name)
	if err != nil {
		return customer.Customer{}, err
	}

	return repo.InsertCustomer(newCustomer)
}
