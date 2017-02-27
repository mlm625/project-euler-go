/*
Solution for Problem 5

Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

Find LCM using greatest common divisor (GCD).
https://en.wikipedia.org/wiki/Least_common_multiple#Reduction_by_the_greatest_common_divisor
 */
package main

// GCD returns the greatest common divisor of two numbers using the Euclidean algorithm.
// See https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a int, b int) int {
	if b == 0 { return a}
	return GCD(b, a % b)
}

// LCM returns the least common multiplier for two numbers using reduction by GCD
// See https://en.wikipedia.org/wiki/Least_common_multiple#Reduction_by_the_greatest_common_divisor
func LCM(a int, b int) int {
	return a / GCD(a, b) * b
}

// SolveGCD gets the LCM for numbers 1 to 20 using the associative property,
// 		LCM(a, b, c) = LCM(LCM(a, b), c)
// See https://en.wikipedia.org/wiki/Least_common_multiple#Lattice-theoretic
func SolveGCD() int {
	lcm := 1
	for i:=1; i<LIMIT; i++ {
		lcm = LCM(lcm, i)
	}

	return lcm
}

