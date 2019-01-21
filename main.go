package main

import (
	"fmt"
)

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	// m1 := Matrix{
	// 	Vector{3, 11, 1, 2, 3},
	// 	Vector{2, 4, 4, 5, 6},
	// 	Vector{1, 3, 7, 12, 32},
	// 	Vector{4, -2, 44, 32, 22},
	// 	Vector{5, -1, 0, 0, 1},
	// }

	m2 := Matrix{
		Vector{3, 0, 2},
		Vector{2, 0, -2},
		Vector{0, 1, 1},
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

	a, _ := m2.MinorsArray()
	for _, x := range a {
		fmt.Println(x)
	}
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
