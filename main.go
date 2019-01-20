package main

import "fmt"

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	// m1 := Matrix{
	// 	Vector{1, 2},
	// 	Vector{2, 4},
	// 	Vector{5, 1},
	// }

	// m2 := Matrix{
	// 	Vector{1, 2, -3, 0},
	// 	Vector{2, 4, 1, 0},
	// 	Vector{-5, 1, 2, 0},
	// 	Vector{6, 12, 0, -5},
	// }

	m3 := Matrix{
		Vector{5, -2, 22},
		Vector{8, 89, 5},
		Vector{66, 112, 52},
	}

	// a, err := m1.Add(m2)
	// checker(err)
	// fmt.Println(a)

	// fmt.Println(m1.Negative())
	// fmt.Println(m2.Negative())
	fmt.Println(m3)
	fmt.Println()
	a, err := m3.Determinant()
	checker(err)
	fmt.Println(a)
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
