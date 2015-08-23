package euler

import (
	"testing"
)

func TestProblem001(t *testing.T) {
	expected := uint(233168)
	got := problem001()
	if got != expected {
		t.Errorf("Sum of all multiples of 3 or 5 below 1000. Exepcted %v got %v", expected, got)
	}
}

func TestProblem002(t *testing.T) {
	expected := uint(4613732)
	got := problem002()
	if got != expected {
		t.Errorf("Expected %v Got %v", expected, got)
	}
}

func TestProblem004(t *testing.T) {
	expected := 906609
	got := problem004()
	if got != expected {
		t.Error(expected, got)
	}
}

func TestProblem005(t *testing.T) {
	expected := 232792560
	got := problem005()
	if got != expected {
		t.Error(expected, got)
	}
}

func TestProblem006(t *testing.T) {
	expected := 25164150
	got := problem006()
	if expected != got {
		t.Error(expected, got)
	}
}

func TestProblem007(t *testing.T) {
	expected := 104743
	got := problem007()
	if expected != got {
		t.Error(expected, got)
	}
}
