package main

import "math"

// QuadSolve x
func QuadSolve(a, b, c float64) ([]float64, []complex128) {
	// warning Complex Solve has a problem
	var solution []float64
	var isolution []complex128

	dis := math.Sqrt(math.Pow(b, 2) - 4*a*c)

	if math.Pow(b, 2)-4*a*c < 0 {
		p1 := math.Sqrt(math.Pow(b, 2))
		p2 := math.Sqrt(4 * a * c)
		isolution = append(isolution, complex(p1/(2*a), p2/(2*a)))
		isolution = append(isolution, complex(p1/(2*a), -1*p2/(2*a)))
		return nil, isolution
	}

	solution = append(solution, ((-1*b)+dis)/(2*a))
	solution = append(solution, ((-1*b)-dis)/(2*a))

	return solution, nil
}
