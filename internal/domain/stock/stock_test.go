package stock

import (
	"fmt"
	"testing"
)

func TestNewStock_Valid(t *testing.T) {
	stock, err := NewStock(1, 101, 50)
	if err != nil {
		t.Fatalf("expected valid stock, got error: %v", err)
	}

	if stock.ID() != 1 {
		t.Errorf("expected ID 1, got %d", stock.ID())
	}

	if stock.ProductID() != 101 {
		t.Errorf("expected ProductID 101, got %d", stock.ProductID())
	}

	if stock.Quantity() != 50 {
		t.Errorf("expected Quantity 50, got %d", stock.Quantity())
	}
}

func TestNewStock_InvalidID(t *testing.T) {
	_, err := NewStock(-1, 101, 10)
	if err == nil {
		t.Error("expected error for negative stock ID")
	}
}

func TestNewStock_InvalidProductID(t *testing.T) {
	_, err := NewStock(1, -5, 10)
	if err == nil {
		t.Error("expected error for negative Product ID")
	}
}

func TestNewStock_NegativeQuantity(t *testing.T) {
	_, err := NewStock(1, 101, -10)
	if err == nil {
		t.Error("expected error for negative Quantity")
	}
}

func TestStock_String(t *testing.T) {
	stock, _ := NewStock(1, 101, 20)
	expected := fmt.Sprintf("{'ID': '%d', 'product_id': '%d', 'Quantity': %d}", 1, 101, 20)
	if stock.String() != expected {
		t.Errorf("expected %q, got %q", expected, stock.String())
	}
}
