package main

import (
	"fmt"
)

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	m1 := Matrix{
		Vector{1, 0},
		Vector{0, 1},
	}

	// m2 := Matrix{
	// 	Vector{1, 2},
	// 	Vector{3, 4},
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

	// a, _ := m1.Divide(m2)
	// for _, x := range a {
	// 	fmt.Println(x)
	// }
	fmt.Println(IsOrthogonal(m1))
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
