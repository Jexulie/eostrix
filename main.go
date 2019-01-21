package main

import (
	"fmt"
)

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	m1 := Matrix{
		Vector{1, 2, 3, 5, 6},
		Vector{7, 8, 9, 2, 3},
		Vector{1, 0, 4, 6, 5},
		Vector{4, 3, 5, 5, 6},
		Vector{5, 7, 5, 7, 8},
		Vector{0, 9, 8, 3, 4},
		Vector{5, 2, 2, 1, 4},
		Vector{8, 7, 6, 5, 4},
		Vector{3, 2, 1, 1, 4},
		Vector{5, 6, 7, 8, 3},
	}

	m2 := Matrix{
		Vector{-3},
		Vector{2},
		Vector{5},
		Vector{-7},
		Vector{9},
	}

	// m3 := Matrix{
	// 	Vector{1, 0, 0, 0},
	// 	Vector{0, 1, 0, 0},
	// 	Vector{0, 0, 1, 0},
	// 	Vector{0, 0, 0, 1},
	// }

	// a, err := m1.Add(m2)
	// checker(err)
	// fmt.Println(a)

	// fmt.Println(m1.Negative())
	// fmt.Println(m2.Negative())

	a, _ := m1.Multiply(m2)
	for _, x := range a {
		fmt.Println(x)
	}
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
