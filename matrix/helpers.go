package matrix

import (
	"errors"
	"fmt"
)

//! Helpers

// GetLen x
func (v Vector) GetLen() int {
	return len(v)
}

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

// GetColumnsCount returns list of all columns
func (m Matrix) GetColumnsCount() []int {
	var lengths []int
	for _, v := range m {
		lengths = append(lengths, v.GetLen())
	}
	return lengths
}

// GetColumnCount returns only one column length
func (m Matrix) GetColumnCount() int {
	var lengths []int
	for _, v := range m {
		lengths = append(lengths, v.GetLen())
	}
	return lengths[0]
}

// GetRowCount x
func (m Matrix) GetRowCount() int {
	rows := 0
	for range m {
		rows++
	}
	return rows
}

// CheckColumnCounts x
func (m Matrix) CheckColumnCounts() bool {
	columns := m.GetColumnsCount()

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
func IsSameDimensions(m1, m2 Matrix) bool {
	if m1.GetRowCount() != m2.GetRowCount() {
		return false
	}

	l1 := m1.GetColumnsCount()
	l2 := m2.GetColumnsCount()

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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	rows := m.GetRowCount()
	dim := m[0].GetLen()
	for _, v := range m {
		l := v.GetLen()
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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return false, errors.New("matrix is not square")
	}
	for i := 0; i < m.GetRowCount(); i++ {
		for j := 0; j < m[i].GetLen(); j++ {
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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return false, errors.New("matrix is not square")
	}
	for i := 0; i < m.GetRowCount(); i++ {
		for j := 0; j < m[i].GetLen(); j++ {
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

// GetSubMatrix x
// maybe check for squareness of matrix
func (m Matrix) GetSubMatrix() []Matrix {
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	var subMatrices []Matrix
	var newMatrix Matrix

	rows := m.GetRowCount()
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
			if newVector.GetLen() != 0 {
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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() != true {
		return nil, errors.New("matrix is not square")
	}
	if m.GetRowCount() == 2 {
		// maybe Manuel Fix for now!
		return Matrix{
			Vector{m[1][1], m[0][1] * -1},
			Vector{m[1][0] * -1, m[0][0]},
		}, nil

	} else if m.GetRowCount() >= 3 {
		// TODO for 3x3 or more
	}
	return nil, nil
}

// GetCofactorMatrix a.k.a GetCofactorMatrix
func (m Matrix) GetCofactorMatrix() (Matrix, error) {
	if m.IsSquare() == false {
		return nil, errors.New("matrix is not square")
	}
	if m.GetRowCount() < 3 && m.GetColumnCount() < 3 {
		return m, nil
	}
	var subMatrices []Matrix
	var newMatrix Matrix

	rows := m.GetRowCount()
	column := m.GetColumnCount()

	// future change this with subMatrix
	for y := 0; y < column; y++ {
		for z := 0; z < rows; z++ {
			for i, v := range m {
				var newVector Vector
				for j := range v {
					if i != y {
						if j != z {
							newVector = append(newVector, m[i][j])
						}
					}
				}
				if newVector.GetLen() != 0 {
					newMatrix = append(newMatrix, newVector)
				}
			}
			subMatrices = append(subMatrices, newMatrix)
			newMatrix = nil
		}
	}

	var minorMatrix Matrix

	guide := len(subMatrices)
	g := 0

	for x1 := 0; x1 < rows; x1++ {
		var minorVector Vector
		if g > guide {
			break //maybe Safeguard
		}
		for y1 := 0; y1 < column; y1++ {
			d, _ := subMatrices[g].Determinant()
			minorVector = append(minorVector, d)
			g++
		}
		minorMatrix = append(minorMatrix, minorVector)
	}

	return minorMatrix, nil
}

// IsOrthogonal x
func IsOrthogonal(m Matrix) (bool, error) {
	if m.IsSquare() == false {
		return false, errors.New("matrix is not square")
	}
	trans := m.Transpose()
	inverse, err := m.Inverse()

	if err != nil {
		return false, err
	}

	for i := 0; i < trans.GetRowCount(); i++ {
		for j := 0; j < trans.GetColumnCount(); j++ {
			if trans[i][j] != inverse[i][j] {
				return false, nil
			}
		}
	}

	return true, nil
}
