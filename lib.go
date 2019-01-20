package main

import (
	"errors"
)

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

// Negative x
func (m Matrix) Negative() Matrix {
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
			newVector = append(newVector, m[i][j]-m2[i][j])
		}
		newMatrix = append(newMatrix, newVector)
	}

	return newMatrix, nil
}

// Divide x
func (m Matrix) Divide(m2 Matrix) (Matrix, error) {
	var newMatrix Matrix
	// need inverse for this
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

// GetSubMatrix x // maybe check for squareness of matrix
func (m Matrix) GetSubMatrix() []Matrix {

	var subMatrices []Matrix
	var newMatrix Matrix

	rows := m.getRowCount()
	for z := 0; z < rows; z++ {
		for i, v := range m {
			var newVector Vector
			for j := range v {
				if i != 0 {
					if j != z {
						newVector = append(newVector, m[i][j])
					}
				}
			}
			if newVector.getLen() != 0 {
				newMatrix = append(newMatrix, newVector)
			}
		}
		subMatrices = append(subMatrices, newMatrix)
		newMatrix = nil
	}
	return subMatrices
}

// Determinant x
func (m Matrix) Determinant() (float64, error) {
	s := m.IsSquare()
	if s != true {
		return 0, errors.New("matrix is not square")
	}
	if m.getRowCount() == 2 {
		return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0]), nil
	} else if m.getRowCount() >= 3 {
		det := 0.0
		subMatrices := m.GetSubMatrix()
		for i, v := range m[0] {
			d, err := subMatrices[i].Determinant()
			if err != nil {
				return 0, err
			}
			if i%2 == 0 {
				det += v * d
			} else {
				det += v * -1 * d
			}
		}
		return det, nil
	} else {
		return 0, errors.New("unknown error")
	}
}

// ! old 3-dim matrix det solution
// } else if m.getRowCount() == 3 {
// 	fPart := Matrix{
// 		Vector{m[1][1], m[1][2]},
// 		Vector{m[2][1], m[2][2]},
// 	}
// 	sPart := Matrix{
// 		Vector{m[1][0], m[1][2]},
// 		Vector{m[2][0], m[2][2]},
// 	}
// 	tPart := Matrix{
// 		Vector{m[1][0], m[1][1]},
// 		Vector{m[2][0], m[2][1]},
// 	}

// 	detF, _ := fPart.Determinant()
// 	detS, _ := sPart.Determinant()
// 	detT, _ := tPart.Determinant()

// 	first := detF * m[0][0]
// 	second := detS * m[0][1]
// 	third := detT * m[0][2]

// 	return first - second + third, nil
