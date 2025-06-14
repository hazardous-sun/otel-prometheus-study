package shared

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewPrice_ValidInputs(t *testing.T) {
	cases := []struct {
		input         string
		expectedValue string
	}{
		{"0.00", "0.00"},
		{"19.99", "19.99"},
		{"100", "100.00"},
		{"0.01", "0.01"},
	}

	for _, tc := range cases {
		price, err := NewPrice(tc.input)
		if err != nil {
			t.Errorf("expected no error for input %q, got %v", tc.input, err)
			continue
		}

		expected, _ := decimal.NewFromString(tc.expectedValue)
		if !price.Value().Equal(expected) {
			t.Errorf("expected value %s, got %s", expected, price.Value())
		}

		if price.String() != expected.StringFixed(2) {
			t.Errorf("expected string %s, got %s", expected.StringFixed(2), price.String())
		}
	}
}

func TestNewPrice_InvalidInputs(t *testing.T) {
	cases := []string{
		"abc",
		"",
		"--12",
		"19,99", // comma instead of dot
		"   ",
	}

	for _, input := range cases {
		_, err := NewPrice(input)
		if err == nil {
			t.Errorf("expected error for invalid input %q, got nil", input)
		}
	}
}

func TestNewPrice_NegativeInput(t *testing.T) {
	_, err := NewPrice("-1.00")
	if err == nil {
		t.Errorf("expected error for negative price, got nil")
	}
}
