/*
Solution for problem 3

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
package main

import (
	"fmt"
	"github.com/go-project-euler/util"
)

const LIMIT = 600851475143

func Solve() int {
	return util.GetLPF(LIMIT)
}

func main() {
	fmt.Printf("Largest Prime Factor: %v\n", Solve())
}
