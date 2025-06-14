package shared

import (
	"testing"
)

func TestNewID_ValidIDs(t *testing.T) {
	ids := []int{0, 1, 213, 5215, 2123, 6465, 1000, 25413, 54212, 653234623}

	for _, input := range ids {
		id, err := NewID(input)
		if err != nil {
			t.Errorf("expected valid ID for input %d, got error: %v", input, err)
		}
		if id.Value() != input {
			t.Errorf("expected id.Value() to equal %d, got %d", input, id.Value())
		}
	}
}

func TestNewID_InvalidIDs(t *testing.T) {
	ids := []int{-1, -213, -5215, -2123, -6465, -1000, -25413, -54212, -653234623}

	for _, input := range ids {
		_, err := NewID(input)
		if err == nil {
			t.Errorf("expected error for invalid ID input %d, but got none", input)
		}
	}
}
