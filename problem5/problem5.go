/*
Solution for Problem 5

Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

Find LCM using prime factors.
https://en.wikipedia.org/wiki/Least_common_multiple#Finding_least_common_multiples_by_prime_factorization
 */
package main

import (
	"fmt"
	"math"
)

const LIMIT = 20
const PRIMES=0
const OCCUR=1

type PrimeTable [2][8] int

func getPrimeExponent(num int, prime int) int {
	count := 0
	for ; num % prime == 0; num /= prime { count += 1 }
	return count
}

func Solve() int {
	pt := PrimeTable{
		[8]int{2, 3, 5, 7, 11, 13, 17, 19},
		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
	}

	// find greatest prime factor exponent for each number
	for i:=1; i <= LIMIT; i++ {
		for j, num:=0, i; j < len(pt[PRIMES]) && num > 1; j++ {
			count := getPrimeExponent(num, pt[PRIMES][j])
			if count > pt[OCCUR][j] {pt[OCCUR][j] = count}
		}
	}

	// Get LCM by finding product of all prime factors to the greatest exponent.
	lcm := 1
	for i:=0; i< len(pt[PRIMES]); i++ {
		lcm *= int(math.Pow(float64(pt[PRIMES][i]), float64(pt[OCCUR][i])))
	}

	return lcm
}

func main() {
	fmt.Printf("Smallest multiple: %v\n", Solve())
}
