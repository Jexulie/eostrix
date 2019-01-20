package main

import "fmt"

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	// m1 := Matrix{
	// 	Vector{3, 4},
	// 	Vector{6, 8},
	// }

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
	for j := 0; j < 10; j++ {
		a, err := RandomMatrix(2, 2, 25)
		checker(err)
		fmt.Println(a)
		// for _, i := range a {
		// 	fmt.Println(i)
		// }
		fmt.Println()
		fmt.Println(a.Determinant())
	}
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
