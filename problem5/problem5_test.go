package main


import "testing"

func TestSolve(t *testing.T) {
	expected := 232792560
	answer := Solve()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}
}

func TestSolveGCD(t *testing.T) {
	expected := 232792560
	answer := SolveGCD()
	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}

}
