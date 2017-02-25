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
	for i := 999; i >= 500; i-- {
		for j := i; j >= 500; j-- {
			prod := i * j
			if isPalindrome(prod) && lpd < prod {
				lpd = prod
			}
		}
	}

	return lpd
}

func main() {
	fmt.Printf("Largest Palindrome Product: %v\n", Solve())
}
