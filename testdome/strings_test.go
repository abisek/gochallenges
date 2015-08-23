package testdome

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		arg1, arg2 string
		expected   bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"b", "a", false},
		{"gopher", "rephog", true},
		{"gopher", "RepHog", false},
		{"gopher", "gophers", false},
	}

	for _, test := range tests {
		actual := IsAnagram(test.arg1, test.arg2)
		if test.expected != actual {
			t.Error(test, actual)
		}
	}

}
