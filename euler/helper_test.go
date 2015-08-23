package euler

import (
	"testing"
)

func TestReverse(t *testing.T) {
	inputs := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"  ", "  "},
		{"abc", "cba"},
	}

	for _, test := range inputs {
		actual := reverse(test.input)
		if test.expected != actual {
			t.Error(test, actual)
		}
	}
}

func TestPalindrome(t *testing.T) {
	inputs := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"abab", false},
	}

	for _, test := range inputs {
		actual := isPalindrome(test.input)
		if actual != test.expected {
			t.Error(test, actual)
		}
	}
}
