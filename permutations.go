package main

import (
	"math"
)

// Perm x
func Perm(totalNumber float64, choice float64, repetition bool) float64 {
	if repetition == true {
		return math.Pow(totalNumber, choice)
	}

	return Factorial(totalNumber) / Factorial((totalNumber - choice))
}

// Permutations x
func Permutations(list []int, c int) [][]int {
	var result [][]int
	//! :(
	return result
}
