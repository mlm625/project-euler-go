/*
Solution for Problem 1

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6
and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

 */

package main

import "fmt"

const LIMIT = 1000

func Solve() int {
	sum := 0

	for i := 3; i < LIMIT; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Printf("Sum: %v\n", Solve())
}


