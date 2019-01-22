package main

import "fmt"

func main() {
	s := Pressure{105, "Bar"}
	s.ToATM()
	fmt.Println(s)
}
func checker(err error) {
	if err != nil {
		panic(err)
	}
}
