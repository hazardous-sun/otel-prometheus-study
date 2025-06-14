package customer

import (
	"fmt"
	"testing"
)

func TestNewCustomer_ValidCustomer(t *testing.T) {
	ids := []int{0, 1, 101, 202, 999, 1234, 4567, 8910, 11213, 1415}
	names := []string{"John Doe", "Alice Smith", "Bob Johnson", "Mary Jane", "Clark Kent", "Bruce Wayne", "Peter Parker", "Diana Prince", "Tony Stark", "Steve Rogers"}

	for i := range ids {
		customer, err := NewCustomer(ids[i], names[i])
		if err != nil {
			t.Errorf("expected valid customer for input %d, got error: %v", ids[i], err)
		}
		if customer.ID() != ids[i] {
			t.Errorf("expected customer.ID() to equal %d, got %d", ids[i], customer.ID())
		}
		if customer.Name() != names[i] {
			t.Errorf("expected customer.Name() to equal %q, got %q", names[i], customer.Name())
		}
	}
}

func TestNewCustomer_InvalidID(t *testing.T) {
	ids := []int{-1, -20, -300, -4000, -50000}
	names := []string{"John Doe", "Alice Smith", "Bob Johnson", "Mary Jane", "Clark Kent"}

	for i := range ids {
		_, err := NewCustomer(ids[i], names[i])
		if err == nil {
			t.Errorf("expected invalid customer ID %d to generate error", ids[i])
		}
	}
}

func TestNewCustomer_InvalidName(t *testing.T) {
	ids := []int{1, 2, 3, 4, 5}
	names := []string{"", "Al!ce", "B0b", "Mar#y", "1234"}

	for i := range ids {
		_, err := NewCustomer(ids[i], names[i])
		if err == nil {
			t.Errorf("expected invalid customer name %q to generate error", names[i])
		}
	}
}

func TestCustomer_String(t *testing.T) {
	id := 100
	name := "Sample Customer"

	customer, err := NewCustomer(id, name)
	if err != nil {
		t.Fatalf("expected valid customer, got error: %v", err)
	}

	expectedString := fmt.Sprintf("{'id': '%d', 'name': '%s'}", id, name)
	if customer.String() != expectedString {
		t.Errorf("expected customer.String() to return %q, got %q", expectedString, customer.String())
	}
}
