package shared

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const (
	nameMinLength = 3
	nameMaxLength = 255
)

var (
	// Regex for name validation
	// Supports Unicode letters, spaces, hyphens, apostrophes
	validNameRegex = regexp.MustCompile(`^[\p{L} '-]+$`)
)

// Name
// Represents a validated name.
type Name struct {
	value string
}

func (n Name) Value() string {
	return n.value
}

func (n Name) String() string {
	return n.value
}

// NewName
// Validates and constructs a new Name.
// Rejects names with digits, special characters, or control characters.
func NewName(value string) (Name, error) {
	trimmed := strings.TrimSpace(value)

	// Check for minimum name length
	if len(trimmed) < nameMinLength {
		return Name{}, errors.New(fmt.Sprintf("NewName(): name must be at least %d characters long", nameMinLength))
	}

	// Check for maximum name length
	if len(trimmed) > nameMaxLength {
		return Name{}, errors.New(fmt.Sprintf("NewName(): name must be at most %d characters long", nameMaxLength))
	}

	// Check if the name contains special characters or numbers
	if !validNameRegex.MatchString(trimmed) {
		return Name{}, errors.New("NewName(): name must contain only letters")
	}

	// Check for control characters
	for _, r := range trimmed {
		if unicode.IsControl(r) {
			return Name{}, errors.New("NewName(): name must not contain control characters")
		}
	}

	return Name{value: trimmed}, nil
}
