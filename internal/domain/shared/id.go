package shared

import "fmt"

type ID struct {
	value int
}

func (i ID) Value() int {
	return i.value
}

func (i ID) String() string {
	return string(rune(i.value))
}

func NewID(value int) (ID, error) {
	// Check if the value is a valid ID
	if value < 0 {
		return ID{}, fmt.Errorf("NewID(): ID should be equal or greater than zero")
	}

	return ID{value: value}, nil
}
