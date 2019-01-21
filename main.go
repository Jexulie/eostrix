package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	x := Permutations(a, 2)
	for _, v := range x {
		fmt.Println(v)
	}
}
func checker(err error) {
	if err != nil {
		panic(err)
	}
}
