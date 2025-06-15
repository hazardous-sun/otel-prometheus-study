package store_product

import (
	"fmt"
	"testing"
)

func TestNewStoreProduct_Valid(t *testing.T) {
	sp, err := NewStoreProduct(1, 100, "19.99", 10)
	if err != nil {
		t.Fatalf("expected valid StoreProduct, got error: %v", err)
	}

	if sp.StoreID() != 1 || sp.ProductID() != 100 || sp.Price() != "19.99" || sp.Quantity() != 10 {
		t.Errorf("unexpected values in StoreProduct: %+v", sp)
	}
}

func TestNewStoreProduct_InvalidIDs(t *testing.T) {
	_, err := NewStoreProduct(-1, 100, "10.00", 5)
	if err == nil {
		t.Error("expected error for negative StoreID")
	}
	_, err = NewStoreProduct(1, -5, "10.00", 5)
	if err == nil {
		t.Error("expected error for negative ProductID")
	}
}

func TestNewStoreProduct_InvalidPrice(t *testing.T) {
	_, err := NewStoreProduct(1, 1, "abc", 5)
	if err == nil {
		t.Error("expected error for invalid Price")
	}
}

func TestNewStoreProduct_NegativeQuantity(t *testing.T) {
	_, err := NewStoreProduct(1, 1, "9.99", -5)
	if err == nil {
		t.Error("expected error for negative Quantity")
	}
}

func TestStoreProduct_String(t *testing.T) {
	sp, _ := NewStoreProduct(2, 50, "14.00", 25)
	expected := fmt.Sprintf("{'store_id': '%d', 'product_id': '%d', 'Price': '%s', 'Quantity': %d}", 2, 50, "14.00", 25)
	if sp.String() != expected {
		t.Errorf("expected %q, got %q", expected, sp.String())
	}
}
