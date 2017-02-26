/*
Solution for Problem 5

Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
package main

import (
	"fmt"
	"math"
)

const LIMIT = 20

// Find LCM using prime factors. See https://en.wikipedia.org/wiki/Least_common_multiple
func Solve() int {
	primes := []int {2, 3, 5, 7, 11, 13, 17, 19}
	var occurrences [8]int

	// find greatest prime factor exponent for each number
	for i:=1; i <= LIMIT; i++ {
		for j, num:=0, i; j < len(primes) && num > 1; j++ {
			count := 0
			for ; num % primes[j] == 0; num /= primes[j] {
				count += 1
			}
			if count > occurrences[j] {occurrences[j] = count}
		}
	}

	// Get LCM by finding product of all prime factors to the greatest exponent.
	lcm := 1
	for i:=0; i< len(primes); i++ {
		lcm *= int(math.Pow(float64(primes[i]), float64(occurrences[i])))
	}

	return lcm
}

func main() {
	fmt.Printf("Smallest multiple: %v\n", Solve())
}
