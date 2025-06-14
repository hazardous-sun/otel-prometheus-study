package shared

import "fmt"

// ID
// Represents a validated identifier.
type ID struct {
	value int
}

func (i ID) Value() int {
	return i.value
}

func (i ID) String() string {
	return string(rune(i.value))
}

// NewID
// Validates and creates a new ID. It must be non-negative.
func NewID(value int) (ID, error) {
	// Check if the value is a valid ID
	if value < 0 {
		return ID{}, fmt.Errorf("NewID(): ID should be equal or greater than zero")
	}

	return ID{value: value}, nil
}
