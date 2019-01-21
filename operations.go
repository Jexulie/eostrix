package main

import (
	"errors"
)

// TODO Better CheckColumnCounts Error
// maybe return error or panic

// Determinant x
func (m Matrix) Determinant() (float64, error) {
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() != true {
		return 0, errors.New("matrix is not square")
	}
	if m.GetRowCount() == 2 {
		return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0]), nil
	} else if m.GetRowCount() >= 3 {
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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return nil, errors.New("matrix is not square")
	}

	if m.GetRowCount() == 2 {
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
	} else if m.GetRowCount() >= 3 {
		detOriginal, _ := m.Determinant()
		if detOriginal == 0 {
			// panic("Division by Zero Error")
			return nil, errors.New("This Matrix has no Inverse, determinant is zero")
		}
		minors, _ := m.GetCofactorMatrix()
		c := minors.GetColumnCount()
		r := minors.GetRowCount()

		var cofactors Matrix
		for i := 0; i < c; i++ {
			var tempVector Vector
			for j := 0; j < r; j++ {
				n := minors[i][j] * -1
				if (j%2 == 0) == (i%2 == 0) || (j%2 == 1) == (i%2 == 1) {
					n = n * -1
				}
				tempVector = append(tempVector, n)
			}
			cofactors = append(cofactors, tempVector)
		}

		trans := cofactors.Transpose()
		finally := trans.ScalarMulti(1 / detOriginal)
		return finally, nil
	}
	return nil, nil
}

// Transpose x
func (m Matrix) Transpose() Matrix {
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	columns := m.GetColumnCount()
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
	if m.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m.IsSquare() == false {
		return 0.0, errors.New("matrix is not square")
	}
	trace := 0.0
	for i := 0; i < m.GetRowCount(); i++ {
		for j := 0; j < m[i].GetLen(); j++ {
			if i == j {
				trace += m[i][j]
			}
		}
	}
	return trace, nil
}

// Eigenvalues x // future Too hard for now, need quadratic equation solver
func (m Matrix) Eigenvalues() []float64 {
	var values []float64
	return values
}

// Eigenvectors x //future Too hard for now, need quadratic equation solver
func (m Matrix) Eigenvectors() []Vector {
	var vectors []Vector
	return vectors
}
