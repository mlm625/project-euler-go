package main

import "testing"

func TestSolve(t *testing.T) {
	expected := 104743
	answer := Solve()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}
}

func TestSolve2(t *testing.T) {
	expected := 104743
	answer := Solve2()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}
}
