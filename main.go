package main

import "fmt"

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	m1 := Matrix{
		Vector{1, 0, 0, 0},
		Vector{0, 1, 0, 0},
		Vector{0, 0, 1, 0},
		Vector{0, 0, 0},
	}

	// m2 := Matrix{
	// 	Vector{13},
	// 	Vector{8},
	// 	Vector{6},
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
	fmt.Println(m1.Trace())
	// fmt.Println(m2.Trace())
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
