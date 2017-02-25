/*
Solution for problem 3

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
package main

import "fmt"

const LIMIT = 600851475143

// getNextNum gets the next number in a selected formula for Dirichlet's arithmetic progression. These formulas are
// used to generate prime numbers.
// Note: not all numbers generated are primes, but all prime numbers can be generated to infinity.
// See https://en.wikipedia.org/wiki/Dirichlet's_theorem_on_arithmetic_progressions
func getNextNum(n int) int {
	return 2*n + 1
}

func Solve() int {
	lpf, num := 0, LIMIT

	// This is a small number, just use trial division and getNextNum as prime sieve. Composite factors will have
	// been eliminated by preceding prime factors.
	for factor, n := 2, 1; num > 1 ; factor, n = getNextNum(n), n+1 {
		if num % factor == 0 {
			lpf = factor
			for num % factor == 0 { num /= factor }
		}
	}

	return lpf
}

func main() {
	fmt.Printf("Largest Prime Factor: %v\n", Solve())
}
