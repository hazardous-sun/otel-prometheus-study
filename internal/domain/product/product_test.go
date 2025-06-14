package product

import (
	"strconv"
	"testing"
)

func TestNewProduct_ValidProduct(t *testing.T) {
	ids := []int{0, 1, 323, 6115, 27123, 8765, 3000, 76534, 72, 987654321}
	names := []string{"Jack Sparrow", "Sponge Bob", "Goku", "Bob Marley", "Frank Sinatra", "Luffy", "Gandalf The White", "Sauron", "Harry Potter", "Ivankov"}
	prices := []string{"19.99", "42.14", "50", "10000", "22.15", "53.12", "1425.73", "33.3", "3.14", "240.00"}

	for i := range ids {
		product, err := NewProduct(ids[i], names[i], prices[i])
		if err != nil {
			t.Errorf("expected valid product for input %d, got error: %v", ids[i], err)
		}
		if product.Name() != names[i] {
			t.Errorf("expected product.Name() to equal %q, got %q", names[i], product.Name())
		}

		// Compare prices as float64
		expectedPrice, err := strconv.ParseFloat(prices[i], 64)
		if err != nil {
			t.Fatalf("failed to parse expected price %s: %v", prices[i], err)
		}
		actualPrice, err := strconv.ParseFloat(product.Price(), 64)
		if err != nil {
			t.Errorf("failed to parse actual product price %s: %v", product.Price(), err)
		}

		if expectedPrice != actualPrice {
			t.Errorf("expected product.Price() to equal '%f', got '%f'", expectedPrice, actualPrice)
		}
	}
}

func TestNewProduct_InvalidID(t *testing.T) {
	ids := []int{-1, -2, -323, -6115, -27123, -8765, -3000, -76534, -72, -987654321}
	names := []string{"Jack Sparrow", "Sponge Bob", "Goku", "Bob Marley", "Frank Sinatra", "Luffy", "Gandalf The White", "Sauron", "Harry Potter", "Ivankov"}
	prices := []string{"19.99", "42.14", "50", "10000", "22.15", "53.12", "1425.73", "33.3", "3.14", "240.00"}

	for i := range ids {
		_, err := NewProduct(ids[i], names[i], prices[i])
		if err == nil {
			t.Errorf("expected invalid product ID %q to generate error", ids[i])
		}
	}
}

func TestNewProduct_InvalidName(t *testing.T) {
	ids := []int{0, 1, 323, 6115, 27123, 8765, 3000, 76534, 72, 987654321}
	names := []string{"", "Spong1e Bob", "Gok8", "Bob Ma!rley", "F?rank Sinatra", "Monkey D. Luffy", "G4nd4lf Th3 Wh1t3", "?S4ur0n", "44rry Potter", "I1v@ankov"}
	prices := []string{"19.99", "42.14", "50", "10000", "22.15", "53.12", "1425.73", "33.3", "3.14", "240.00"}

	for i := range ids {
		_, err := NewProduct(ids[i], names[i], prices[i])
		if err == nil {
			t.Errorf("expected invalid product name %q to generate error", names[i])
		}
	}
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	ids := []int{0, 1, 323, 6115, 27123, 8765, 3000, 76534, 72, 987654321}
	names := []string{"Jack Sparrow", "Sponge Bob", "Goku", "Bob Marley", "Frank Sinatra", "Luffy", "Gandalf The White", "Sauron", "Harry Potter", "Ivankov"}
	prices := []string{"-19.99", "-42.14", "-50", "-10000", "-22.15", "-53.12", "-1425.73", "-33.3", "-3.14", "-240.00"}

	for i := range ids {
		_, err := NewProduct(ids[i], names[i], prices[i])
		if err == nil {
			t.Errorf("expected invalid product price %q to generate error", prices[i])
		}
	}
}
