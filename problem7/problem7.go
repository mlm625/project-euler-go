/*
10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10,001st prime number?
 */
package main

import (
	"fmt"
	"github.com/go-project-euler/util"
	"math"
)

const LIMIT = 10001

func isPrimeByLPF(n int) bool {
	return n == util.GetLPF(n)
}

func Solve()int {
	prime := 1
	for num, primes:=1, 0; primes <= LIMIT; num++ {
		if isPrimeByLPF(num) { primes++; prime = num }
	}

	return prime
}

func isPrimeByDiv(n int, primeNumbers []int) bool {
	for _, p := range primeNumbers {
		if n % p == 0 {return false}
		if float64(p) > math.Sqrt(float64(n)){break}
	}
	return true
}

func Solve2()int {
	primes := make([]int, 0, LIMIT)
	primes = append(primes, 2)
	for i:=1; len(primes)<LIMIT; i++ {
		n := 2*i+1
		if isPrimeByDiv(n, primes) {primes = append(primes, n)}
	}
	return primes[len(primes)-1]
}


func main() {
	fmt.Printf("The 10,001'st prime: %v\n", Solve2())
}
