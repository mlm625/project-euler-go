package util

// getNextNum gets the next number in a selected formula for Dirichlet's arithmetic progression. These formulas are
// used to generate prime numbers.
// Note: not all numbers generated are primes, but all prime numbers can be generated to infinity.
// See https://en.wikipedia.org/wiki/Dirichlet's_theorem_on_arithmetic_progressions
// Fancy way of saying odd numbers...
func getNextNum(n int) int {
	return 2*n + 1
}

// GetLPF gets the largest prime factor for a number that's not too large.
func GetLPF(number int) int {
	lpf, num := 1, number

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
