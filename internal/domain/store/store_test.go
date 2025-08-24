package store

import (
	"fmt"
	"otel-prometheus-study/internal/domain/store_product"
	"testing"
)

func TestNewStore_Valid(t *testing.T) {
	store, err := NewStore(1, "Downtown Branch")
	if err != nil {
		t.Fatalf("expected valid store, got error: %v", err)
	}

	if store.ID() != 1 {
		t.Errorf("expected store ID to be 1, got %d", store.ID())
	}

	if store.Name() != "Downtown Branch" {
		t.Errorf("expected store Name to be 'Downtown Branch', got %s", store.Name())
	}
}

func TestNewStore_InvalidID(t *testing.T) {
	_, err := NewStore(-10, "Branch A")
	if err == nil {
		t.Error("expected error for negative store ID")
	}
}

func TestNewStore_InvalidName(t *testing.T) {
	_, err := NewStore(10, "!!Invalid##Name")
	if err == nil {
		t.Error("expected error for invalid store Name")
	}
}

func TestStore_String(t *testing.T) {
	store, _ := NewStore(42, "Main HQ")
	expected := fmt.Sprintf("{'ID': '%d', 'Name': '%s', 'products_count': 0}", 42, "Main HQ")

	if store.String() != expected {
		t.Errorf("expected %q, got %q", expected, store.String())
	}
}

func TestStore_AddProduct(t *testing.T) {
	store, _ := NewStore(1, "Central")
	sp, _ := store_product.NewStoreProduct(1, 101, "15.00", 50)

	store, err := store.AddProduct(sp)
	if err != nil {
		t.Fatalf("expected product to be added successfully, got error: %v", err)
	}

	if len(store.Products()) != 1 {
		t.Errorf("expected 1 product, got %d", len(store.Products()))
	}
}

func TestStore_AddProduct_InvalidStoreID(t *testing.T) {
	store, _ := NewStore(1, "Branch A")
	sp, _ := store_product.NewStoreProduct(2, 105, "12.50", 10) // storeID != store.ID()

	_, err := store.AddProduct(sp)
	if err == nil {
		t.Error("expected error due to mismatched store IDs")
	}
}
