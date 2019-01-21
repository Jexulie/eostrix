package matrix

import (
	"errors"
	"math/rand"
	"time"
)

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

// CreateZero x
func CreateZero(rows, columns int) (Matrix, error) {
	var newMatrix Matrix
	for i := 0; i < rows; i++ {
		var newVector Vector
		for j := 0; j < columns; j++ {
			newVector = append(newVector, 0)
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
				rand.Seed(time.Now().UTC().UnixNano())
				time.Sleep(100 * time.Microsecond)
				fakeSeed *= -1
			}
			newVector = append(newVector, rand.Float64()*float64(fakeSeed))
		}
		newMatrix = append(newMatrix, newVector)
	}
	return newMatrix, nil
}
