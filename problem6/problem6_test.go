package main


import "testing"

func TestSolve(t *testing.T) {
	expected := 25164150
	answer := Solve()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}
}
