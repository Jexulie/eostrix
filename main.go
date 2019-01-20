package main

import "fmt"

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	m1 := Matrix{
		Vector{4, 7},
		Vector{2, 6},
	}

	// m2 := Matrix{
	// 	Vector{1, 0, 0},
	// 	Vector{0, 1, 0},
	// 	Vector{0, 0, 1},
	// }

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
	// fmt.Println(m1.Trace())
	// fmt.Println(m2.Trace())
	fmt.Println(m1.Inverse())
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
