package main

import "errors"

// Multiply (Matrix x Matrix)
func (m Matrix) Multiply(m2 Matrix) (Matrix, error) {
	// future Cheap Fix for Column Counts Todo Later!!
	// TODO - WRONG!!!
	if m.CheckColumnCounts() == false || m2.CheckColumnCounts() == false {
		panic("Matrix Columns Defined Wrong")
	}
	if m2.GetRowCount() != m.GetColumnCount() {
		return nil, errors.New("matrix dimensions are not suitable for dot product")
	}
	firstRowCount := m.GetRowCount()
	secondRowCount := m2.GetRowCount()
	secondColumnCount := m2.GetColumnCount()
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
	var newMatrix Matrix
	// need inverse for this
	return newMatrix, nil
}
