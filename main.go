package main

import (
	"fmt"
)

func main() {
	// g := Data{26, 33, 65, 28, 34, 55, 25, 44, 50, 36, 26, 37, 43, 62, 35, 38, 45, 32, 28, 34}
	g1 := Data{1, 3, 3, 4, 6, 7, 8, 8, 5, 5}
	g2 := Data{1, 2, 5, 6, 7, 8, 9, 1, 2}
	// 1 3 3 4 5 | 6 6 7 8 8
	fmt.Println(g1.Quartiles())
	fmt.Println()
	fmt.Println(g2.Quartiles())
	// fmt.Println(g.PearsonModeSkew())
	// fmt.Println(g.PearsonSkewCoefficients())
}
func checker(err error) {
	if err != nil {
		panic(err)
	}
}
