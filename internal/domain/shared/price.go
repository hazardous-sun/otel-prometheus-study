package shared

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Price struct {
	value decimal.Decimal
}

func (p Price) Value() decimal.Decimal {
	return p.value
}

func (p Price) String() string {
	return p.value.StringFixed(2)
}

func NewPrice(amount string) (Price, error) {
	// Parse the amount as a decimal
	d, err := decimal.NewFromString(amount)

	if err != nil {
		return Price{}, errors.New("NewPrice(): invalid amount")
	}

	// Verify that the amount is greater than or equal to zero
	if d.LessThan(decimal.Zero) {
		return Price{}, errors.New("NewPrice(): amount must be positive")
	}

	return Price{value: d}, nil
}
