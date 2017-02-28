/*
Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is,
	12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
	(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum
is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */
package main

import "fmt"

const LIMIT=100

// Solve using square of sum's rule.
// The square of a sum is equal to the sum of the squares of all the summands plus the sum of all the double products
// of the summands in twos, for example:
// 		(a+b+c)^2=a^2+b^2+c^2+2ab+2bc+2ac
// See http://planetmath.org/squareofsum
func Solve() int {
	diff := 0
	for i:=1; i <= LIMIT; i++  {
		for j:=i; j<= LIMIT; j++ {
			if i == j {continue}
			diff += i*j
		}
	}

	// more efficient to iterate over half of the matrix then multiply by 2
	return diff*2
}

func main() {
	fmt.Printf("Sum square difference: %v\n", Solve())
}
