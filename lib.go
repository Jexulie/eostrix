package main

import "errors"

type Vector []float64

type Matrix []Vector

// GetLen x
func (v Vector) getLen() int {
	return len(v)
}

func (m Matrix) getColumnsCount() []int {
	var lengths []int
	for _, v := range m {
		lengths = append(lengths, v.getLen())
	}
	return lengths
}

// GetRowCount x
func (m Matrix) getRowCount() int {
	rows := 0
	for range m {
		rows++
	}
	return rows
}

// IsSameDimensions z
func isSameDimensions(m1, m2 Matrix) bool {
	if m1.getRowCount() != m2.getRowCount() {
		return false
	}

	l1 := m1.getColumnsCount()
	l2 := m2.getColumnsCount()

	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

// IsSquare x
func (m Matrix) IsSquare() bool {
	rows := m.getRowCount()
	dim := m[0].getLen()
	for _, v := range m {
		l := v.getLen()
		if dim != l {
			return false
		}
		dim = l
	}
	if rows == dim {
		return true
	}
	return false
}

// ScalarMulti x
func (m Matrix) ScalarMulti(num float64) Matrix {
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
	var newMatrix Matrix
	t := isSameDimensions(m, m2)
	if t == false {
		return nil, errors.New("matrices dimensions are not same")
	}

	l := m.getRowCount()

	for i := 0; i < l; i++ {
		l2 := m[i].getLen()
		var newVector Vector
		for j := 0; j < l2; j++ {
			newVector = append(newVector, m[i][j]+m2[i][j])
		}
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix, nil
}

// Inverse x
func (m Matrix) Inverse() (Matrix, error) {
	return nil, nil
}

// Transpose x
func (m Matrix) Transpose() (Matrix, error) {
	return nil, nil
}
