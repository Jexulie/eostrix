package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// TODO Better checkColumnCounts Error
// maybe return error or panic

// Vector type
type Vector []float64

// Matrix type
type Matrix []Vector

// GetLen x
func (v Vector) getLen() int {
	return len(v)
}

//! Helpers
func sumInt(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func sumFloat64(array []float64) float64 {
	sum := 0.0
	for _, v := range array {
		sum += v
	}
	return sum
}
func panicer(msg string) {
	if err := recover(); err != nil {
		fmt.Println(msg)
	}
}

//! Matrix Creators

// CreateIdentity x
func CreateIdentity(rows, columns int) (Matrix, error) {
	if rows != columns {
		return nil, errors.New("Identity Matrix must have same row and column length")
	}
	var newMatrix Matrix
	for i := 0; i < rows; i++ {
		var newVector Vector
		for j := 0; j < columns; j++ {
			if i == j {
				if i == j {
					newVector = append(newVector, 1)
				}
			} else {
				if i != j {
					newVector = append(newVector, 0)
				}
			}
		}
		newMatrix = append(newMatrix, newVector)
	}
	return newMatrix, nil
}

// RandomMatrix x
func RandomMatrix(rows, columns int, maxSeed int) (Matrix, error) {
	var newMatrix Matrix
	for i := 0; i < rows; i++ {
		var newVector Vector
		for j := 0; j < columns; j++ {
			fakeSeed := rand.Intn(maxSeed)
			if fakeSeed%(rand.Intn(3)+1) == 0 {
				fakeSeed *= -1
			}
			newVector = append(newVector, rand.Float64()*float64(fakeSeed))
		}
		newMatrix = append(newMatrix, newVector)
	}
	return newMatrix, nil
}

// getColumnsCount returns list of all columns
func (m Matrix) getColumnsCount() []int {
	var lengths []int
	for _, v := range m {
		lengths = append(lengths, v.getLen())
	}
	return lengths
}

// getColumnCount returns only one column length
func (m Matrix) getColumnCount() int {
	var lengths []int
	for _, v := range m {
		lengths = append(lengths, v.getLen())
	}
	return lengths[0]
}

// GetRowCount x
func (m Matrix) getRowCount() int {
	rows := 0
	for range m {
		rows++
	}
	return rows
}

// checkColumnCounts x
func (m Matrix) checkColumnCounts() bool {
	columns := m.getColumnsCount()

	checker := columns[0]
	for _, c := range columns {
		if c != checker {
			return false
		}
		checker = c
	}
	return true
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
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
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

// IsIdentity x
// warning check Rectangular Diagonal Matrices
func (m Matrix) IsIdentity() (bool, error) {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return false, errors.New("matrix is not square")
	}
	for i := 0; i < m.getRowCount(); i++ {
		for j := 0; j < m[i].getLen(); j++ {
			if i == j {
				if m[i][j] != 1 {
					return false, nil
				}
			} else {
				if m[i][j] != 0 {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

// IsDiagonal x
func (m Matrix) IsDiagonal() (bool, error) {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return false, errors.New("matrix is not square")
	}
	for i := 0; i < m.getRowCount(); i++ {
		for j := 0; j < m[i].getLen(); j++ {
			if i == j {
				if m[i][j] == 0 {
					return false, nil
				}
			} else {
				if m[i][j] != 0 {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

// Multiply (Matrix x Matrix)
func (m Matrix) Multiply(m2 Matrix) (Matrix, error) {
	// future Cheap Fix for Column Counts Todo Later!!
	// TODO - WRONG!!!
	if m.checkColumnCounts() == false || m2.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m2.getRowCount() != m.getColumnCount() {
		return nil, errors.New("matrix dimensions are not suitable for dot product")
	}
	firstRowCount := m.getRowCount()
	secondRowCount := m2.getRowCount()
	secondColumnCount := m2.getColumnCount()
	var newMatrix Matrix

	for fc := 0; fc < secondColumnCount; fc++ {
		sc := 0
		var newVector Vector
		for i := 0; i < firstRowCount; i++ {
			sum := 0.0
			for j := 0; j < secondRowCount; j++ {
				sum += m[sc][j] * m2[j][fc]
			}
			newVector = append(newVector, sum)
		}
		newMatrix = append(newMatrix, newVector)
		if sc >= firstRowCount {
			sc = 0
		}
	}

	return newMatrix, nil
}

// ScalarMulti (Matrix x Scalar)
func (m Matrix) ScalarMulti(num float64) Matrix {
	if m.checkColumnCounts() == false {
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
	if m.checkColumnCounts() == false || m2.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
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
	if m.checkColumnCounts() == false {
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
	if m.checkColumnCounts() == false || m2.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	t := isSameDimensions(m, m2)
	if t == false {
		return nil, errors.New("matrices dimensions are not same")
	}
	var newMatrix Matrix

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
	if m.checkColumnCounts() == false || m2.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	var newMatrix Matrix
	// need inverse for this
	return newMatrix, nil
}

// GetSubMatrix x
// maybe check for squareness of matrix
func (m Matrix) GetSubMatrix() []Matrix {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
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

// Swap x
func (m Matrix) Swap() (Matrix, error) {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() != true {
		return nil, errors.New("matrix is not square")
	}
	if m.getRowCount() == 2 {
		// maybe Manuel Fix for now!
		return Matrix{
			Vector{m[1][1], m[0][1] * -1},
			Vector{m[1][0] * -1, m[0][0]},
		}, nil

	} else if m.getRowCount() >= 3 {
		// TODO for 3x3 or more
	}
	return nil, nil
}

// Determinant x
func (m Matrix) Determinant() (float64, error) {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() != true {
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

// Inverse x
func (m Matrix) Inverse() (Matrix, error) {
	// defer panicer("Matrix Columns Defined Wrong")
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return nil, errors.New("matrix is not square")
	}

	if m.getRowCount() == 2 {
		var newMatrix Matrix
		det, _ := m.Determinant()
		swapped, _ := m.Swap()
		for _, v := range swapped {
			var newVector Vector
			for _, n := range v {
				// defer panicer("This Matrix has no Inverse") //! Better Error Management needed..
				if det == 0 {
					// panic("Division by Zero Error")
					return nil, errors.New("This Matrix has no Inverse, determinant is zero")
				}
				newVector = append(newVector, n/det)
			}
			newMatrix = append(newMatrix, newVector)
		}
		return newMatrix, nil
	} else if m.getRowCount() >= 3 {
		// TODO for 3x3 or more
	}
	return nil, nil
}

// Transpose x
func (m Matrix) Transpose() Matrix {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	columns := m.getColumnCount()
	var newMatrix Matrix
	for i := 0; i < columns; i++ {
		var newVector Vector
		for _, v := range m {
			// fmt.Println(v[i])
			newVector = append(newVector, v[i])
		}
		newMatrix = append(newMatrix, newVector)
	}
	return newMatrix
}

// Trace x
func (m Matrix) Trace() (float64, error) {
	if m.checkColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return 0.0, errors.New("matrix is not square")
	}
	trace := 0.0
	for i := 0; i < m.getRowCount(); i++ {
		for j := 0; j < m[i].getLen(); j++ {
			if i == j {
				trace += m[i][j]
			}
		}
	}
	return trace, nil
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
