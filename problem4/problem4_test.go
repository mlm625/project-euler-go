package main

import "testing"

func TestReverseNum(t *testing.T) {
	reverse := reverseNum(1)
	if reverse != 1 {
		t.Errorf("Exected: 1, Got %v", reverse)
	}

	reverse = reverseNum(10)
	if reverse != 1 {
		t.Errorf("Exected: 1, Got %v", reverse)
	}

	reverse = reverseNum(60)
	if reverse != 6 {
		t.Errorf("Exected: 6, Got %v", reverse)
	}

	reverse = reverseNum(12)
	if reverse != 21 {
		t.Errorf("Exected: 21, Got %v", reverse)
	}

	reverse = reverseNum(100)
	if reverse != 1 {
		t.Errorf("Exected: 1, Got %v", reverse)
	}

	reverse = reverseNum(123)
	if reverse != 321 {
		t.Errorf("Exected: 321, Got %v", reverse)
	}

	reverse = reverseNum(1234)
	if reverse != 4321 {
		t.Errorf("Exected: 4321, Got %v", reverse)
	}

	reverse = reverseNum(1234)
	if reverse != 4321 {
		t.Errorf("Exected: 4321, Got %v", reverse)
	}
}

func TestSolve(t *testing.T) {
	expected := 906609
	answer := Solve()

	if expected != answer {
		t.Errorf("Expected: %v, Got %v", expected, answer)
	}

}
