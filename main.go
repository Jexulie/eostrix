package main

import "fmt"

func main() {
	// v1 := Vector{1, 2, 3, 4, 5, 6, 7}
	m1 := Matrix{
		Vector{1, 2},
		Vector{2, 4},
		Vector{5, 1},
	}

	m2 := Matrix{
		Vector{1, 2},
		Vector{2, 4},
		Vector{5, 1},
	}

	fmt.Println(m1)
	fmt.Println(m2)
	a, err := m1.Add(m2)
	checker(err)
	fmt.Println(a)
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}
