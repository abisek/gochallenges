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
		actual := Reverse(test.input)
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
		actual := IsPalindrome(test.input)
		if actual != test.expected {
			t.Error(test, actual)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{71, true},
		{72, false},
		{2293, true},
		{2294, false},
	}

	for _, test := range tests {
		actual := IsPrime(test.num)
		if test.expected != actual {
			t.Error(test, actual)
		}
	}
}
