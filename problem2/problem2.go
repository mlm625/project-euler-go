/*
Solution for Problem 2

Even Fibonacci numbers

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2,
the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the
even-valued terms.
 */
package main

import "fmt"

const LIMIT = 4000000

// fibonacci takes two terms of the sequence and generates the next term.
func fibonacci(x int, y int) (int, int) {
	return y, x + y
}

func Solve() int {
	sum := 0

	for fibX,fibY :=1, 2; fibY < LIMIT; fibX,fibY = fibonacci(fibX, fibY) {
		if fibY % 2 == 0 {
			sum += fibY
		}
	}

	return sum
}

func main() {
	fmt.Printf("Sum: %v\n", Solve())
}
