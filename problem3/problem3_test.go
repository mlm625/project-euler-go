package main


import "testing"

func TestSolve(t *testing.T) {
	expected := 6857
	answer := Solve()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}
}
