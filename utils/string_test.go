package utils

import (
	"testing"
)

func TestUnderscoreName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"TestUnderscoreName", "test_underscore_name"},
		{"testUnderscoreName", "test_underscore_name"},
		{"testUnderscoreName", ""},
	}

	for _, test := range tests {
		name := UnderscoreName(test.input)
		if test.expected != name {
			t.Errorf("UnderscoreName(%q)=%s, expected:%s", test.input, name, test.expected)
		}
	}
}

func TestCamelName(t *testing.T) {
	tests := []struct {
		expected string
		input    string
	}{
		{"TestUnderscoreName", "test_underscore_name"},
		{"testUnderscoreName", "test_underscore_name"},
		{"testUnderscoreName", ""},
	}

	for _, test := range tests {
		name := CamelName(test.input)
		if test.expected != name {
			t.Errorf("CamelName(%q)=%s, expected:%s", test.input, name, test.expected)
		}
	}
}
