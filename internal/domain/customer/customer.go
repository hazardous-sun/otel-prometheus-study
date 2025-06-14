package customer

import (
	"fmt"
	"otel-prometheus-study/internal/domain/shared"
)

// Customer
// Represents a customer entity with an ID and a Name.
type Customer struct {
	id   shared.ID
	name shared.Name
}

func (c Customer) ID() int {
	return c.id.Value()
}

func (c Customer) Name() string {
	return c.name.Value()
}

func (c Customer) String() string {
	return fmt.Sprintf("{'id': '%d', 'name': '%s'}", c.id.Value(), c.name.Value())
}

// NewCustomer
// Creates a new Customer instance after validating ID and Name inputs.
func NewCustomer(inputId int, inputName string) (Customer, error) {
	// Check if inputId is a valid ID
	id, err := shared.NewID(inputId)
	if err != nil {
		return Customer{}, err
	}

	// Check if inputName is a valid name
	name, err := shared.NewName(inputName)
	if err != nil {
		return Customer{}, err
	}

	return Customer{id, name}, nil
}
