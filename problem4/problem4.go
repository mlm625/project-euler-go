/*
Solution for problem 4

Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import "fmt"

const BASE = 10

func reverseNum(num int) int {
	reverse := 0
	for i := num; i > 0; i /= BASE {
		reverse = reverse * BASE + i % BASE
	}
	return reverse
}

func isPalindrome(num int) bool {
	return num == reverseNum(num)
}

func Solve() int {
	lpd := 0
	for i := 999; i >= 100; i-- {
		// eliminate duplication by setting j to i.
		for j := i; j >= 100; j-- {
			prod := i * j
			// no need to check for a palindrome unless the current prod is greater
			if lpd < prod && isPalindrome(prod) {
				lpd = prod
			}
		}
	}

	return lpd
}

func main() {
	fmt.Printf("Largest Palindrome Product: %v\n", Solve())
}
