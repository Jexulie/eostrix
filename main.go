package main

import (
	"fmt"
)

func main() {
	fmt.Println(QuadSolve(12, 5, 7))
}
func checker(err error) {
	if err != nil {
		panic(err)
	}
}
