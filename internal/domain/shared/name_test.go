package shared

import "testing"

func TestNewName_ValidNames(t *testing.T) {
	validNames := []string{
		"Alice",
		"John Smith",
		"Élise Dupont",
		"O'Connor",
		"Jean-Luc",
	}

	for _, input := range validNames {
		name, err := NewName(input)
		if err != nil {
			t.Errorf("expected valid name for input %q, got error: %v", input, err)
		}
		if name.Value() != input {
			t.Errorf("expected name.Value() to equal %q, got %q", input, name.Value())
		}
	}
}

func TestNewName_InvalidNames(t *testing.T) {
	invalidNames := []string{
		"",          // empty
		"A",         // too short
		"Bob123",    // digits
		"Mary!",     // special characters
		"<script>",  // injection
		"Elon_Musk", // underscore
		"  ",        // only spaces
	}

	for _, input := range invalidNames {
		_, err := NewName(input)
		if err == nil {
			t.Errorf("expected error for invalid name input %q, but got none", input)
		}
	}
}
