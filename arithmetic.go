package main

import (
	"errors"
)

// Multiply (Matrix x Matrix)
func (m Matrix) Multiply(m2 Matrix) (Matrix, error) {
	// future Cheap Fix for Column Counts Todo Later!!
	if m.CheckColumnCounts() == false || m2.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m2.GetRowCount() != m.GetColumnCount() {
		return nil, errors.New("matrix dimensions are not suitable for dot product")
	}
	fRowCount := m.GetRowCount()
	sRowCount := m2.GetRowCount()
	sColumnCount := m2.GetColumnCount()
	var newMatrix Matrix

	for a, j := 0, 0; a < fRowCount; a++ {
		var newVector Vector
		for i := 0; i < sColumnCount; {
			sum := 0.0
			for k := 0; k < sRowCount; k++ {
				sum += m[j][k] * m2[k][i]
			}
			newVector = append(newVector, sum)
			i++
		}
		j++
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix, nil
}

// ScalarMulti (Matrix x Scalar)
func (m Matrix) ScalarMulti(num float64) Matrix {
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	var newMatrix Matrix
	for _, v := range m {
		var newV Vector
		for _, e := range v {
			p := e * num
			newV = append(newV, p)
		}
		newMatrix = append(newMatrix, newV)
	}
	return newMatrix
}

// Add x
func (m Matrix) Add(m2 Matrix) (Matrix, error) {
	if m.CheckColumnCounts() == false || m2.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	var newMatrix Matrix
	t := IsSameDimensions(m, m2)
	if t == false {
		return nil, errors.New("matrices dimensions are not same")
	}

	l := m.GetRowCount()

	for i := 0; i < l; i++ {
		l2 := m[i].GetLen()
		var newVector Vector
		for j := 0; j < l2; j++ {
			newVector = append(newVector, m[i][j]+m2[i][j])
		}
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix, nil
}

// Negative x
func (m Matrix) Negative() Matrix {
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	var newMatrix Matrix

	for _, v := range m {
		var newVector Vector
		for _, e := range v {
			newVector = append(newVector, e*-1)
		}
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix
}

// Subtract x
func (m Matrix) Subtract(m2 Matrix) (Matrix, error) {
	if m.CheckColumnCounts() == false || m2.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	t := IsSameDimensions(m, m2)
	if t == false {
		return nil, errors.New("matrices dimensions are not same")
	}
	var newMatrix Matrix

	l := m.GetRowCount()

	for i := 0; i < l; i++ {
		l2 := m[i].GetLen()
		var newVector Vector
		for j := 0; j < l2; j++ {
			newVector = append(newVector, m[i][j]-m2[i][j])
		}
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix, nil
}

// Divide x
func (m Matrix) Divide(m2 Matrix) (Matrix, error) {
	if m.CheckColumnCounts() == false || m2.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	inv, err := m2.Inverse()
	if err != nil {
		return nil, err
	}

	result, err := m.Multiply(inv)
	if err != nil {
		return nil, err
	}
	return result, nil
}
