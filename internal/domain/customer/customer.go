package customer

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// Customer
// Represents a customer entity with an ID and a Name.
type Customer struct {
	IDValue   shared.ID   `json:"id"`
	NameValue shared.Name `json:"name"`
}

func (c Customer) ID() int {
	return c.IDValue.Value()
}

func (c Customer) Name() string {
	return c.NameValue.Value()
}

func (c Customer) String() string {
	return fmt.Sprintf("{'ID': '%d', 'Name': '%s'}", c.IDValue.Value(), c.NameValue.Value())
}

// NewCustomer
// Creates a new Customer instance after validating ID and Name inputs.
func NewCustomer(inputId int, inputName string) (Customer, error) {
	// Check if inputId is a valid ID
	id, err := shared.NewID(inputId)
	if err != nil {
		return Customer{}, err
	}

	// Check if inputName is a valid Name
	name, err := shared.NewName(inputName)
	if err != nil {
		return Customer{}, err
	}

	return Customer{id, name}, nil
}
