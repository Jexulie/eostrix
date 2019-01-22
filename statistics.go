package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Data x
type Data []float64

// Min x
func (d Data) Min() float64 {
	min := d[0]

	for _, v := range d {
		if v < min {
			min = v
		}
	}

	return min
}

// Max x
func (d Data) Max() float64 {
	max := d[0]

	for _, v := range d {
		if v > max {
			max = v
		}
	}

	return max
}

// Mean x
func (d Data) Mean() float64 {
	return d.Sigma() / float64(len(d))
}

// Mode x
func (d Data) Mode() float64 {
	// future bimodal or multimodal
	m := make(map[string]int)

	includes := func(m map[string]int, val string) bool {
		for k := range m {
			if k == val {
				return true
			}
		}
		return false
	}

	fetchMode := func(m map[string]int) float64 {
		highest := ""
		value := 0
		for k, v := range m {
			if v > value {
				highest = k
				value = v
			}
		}
		h, _ := strconv.ParseFloat(highest, 64)
		return h
	}

	for _, v := range d {
		n := strconv.Itoa(int(v))
		if includes(m, n) {
			m[n]++
		} else {
			m[n] = 1
		}
	}

	return fetchMode(m)
}

// Range x
func (d Data) Range() float64 {
	return math.Abs(d.Max() - d.Min())
}

// Median x
func (d Data) Median() float64 {
	l := len(d)
	if l%2 != 0 {
		return d[l/2]
	}
	return (d[l/2] + d[(l/2)-1]) / 2
}

// Variance x
func (d Data) Variance() float64 {
	l := float64(len(d))
	m := d.Mean()
	s := 0.0
	for _, v := range d {
		s += math.Pow((v - m), 2)
	}
	return (1 / l) * s
}

// SampVariance x
func (d Data) SampVariance() float64 {
	l := float64(len(d))
	m := d.Mean()
	s := 0.0
	for _, v := range d {
		s += math.Pow((v - m), 2)
	}
	return (1 / (l - 1)) * s
}

// Sigma x sum of Data
func (d Data) Sigma() float64 {
	sum := 0.0
	for _, v := range d {
		sum += v
	}
	return sum
}

// Sum x sum of a List
func (d Data) Sum(list []float64) float64 {
	sum := 0.0
	for _, v := range list {
		sum += v
	}
	return sum
}

// Pi x
func (d Data) Pi() float64 {
	prod := 0.0
	for _, v := range d {
		prod *= v
	}
	return prod
}

// Quartiles x
func (d Data) Quartiles() map[string]float64 {
	sort.Float64s(d)
	fmt.Println(d)
	var q1 float64
	var q2 float64
	var q3 float64

	if len(d)%2 == 0 {
		fHalf := d[:len(d)/2]
		sHalf := d[len(d)/2:]
		fmt.Println(fHalf)
		fmt.Println(sHalf)
		q2 = (fHalf[len(sHalf)-1] + sHalf[0]) / 2
		if len(fHalf)%2 != 0 {
			q1 = fHalf[len(fHalf)/2]
		} else {
			m := int(math.Floor(float64(len(fHalf) / 2)))
			n := int(math.Ceil(float64(len(fHalf) / 2)))
			q1 = (fHalf[m-1] + fHalf[n]) / 2
		}
		if len(sHalf)%2 != 0 {
			q3 = sHalf[len(sHalf)/2]
		} else {
			m := int(math.Floor(float64(len(sHalf) / 2)))
			n := int(math.Ceil(float64(len(sHalf) / 2)))
			q3 = (sHalf[m-1] + sHalf[n]) / 2
		}
	} else {
		fHalf := d[:len(d)/2+1]
		sHalf := d[len(d)/2:]
		fmt.Println(fHalf)
		fmt.Println(sHalf)
		q2 = (fHalf[len(sHalf)-1] + sHalf[0]) / 2
		if len(fHalf)%2 != 0 {
			q1 = fHalf[len(fHalf)/2]
		} else {
			m := int(math.Floor(float64(len(fHalf) / 2)))
			n := int(math.Ceil(float64(len(fHalf) / 2)))
			q1 = (fHalf[m-1] + fHalf[n]) / 2
		}
		if len(sHalf)%2 != 0 {
			q3 = sHalf[len(sHalf)/2]
		} else {
			m := int(math.Floor(float64(len(sHalf) / 2)))
			n := int(math.Ceil(float64(len(sHalf) / 2)))
			q3 = (sHalf[m-1] + sHalf[n]) / 2
		}
	}

	return map[string]float64{
		"Q1":     q1,
		"Q2":     q2,
		"Q3":     q3,
		"IRange": q3 - q1,
	}
}

// StdPop x
func (d Data) StdPop() float64 {
	return math.Sqrt(d.Variance())
}

// StdSamp x
func (d Data) StdSamp() float64 {
	return math.Sqrt(d.SampVariance())
}

// IsNormalDist x
func (d Data) IsNormalDist() bool {
	if d.Mean() == d.Median() && d.Mean() == d.Mode() {
		return true
	}
	return false
}

// MeanDeviation x
func (d Data) MeanDeviation() float64 {
	sum := 0.0
	m := d.Mean()
	l := float64(len(d))
	for _, v := range d {
		sum += math.Abs(v - m)
	}
	return sum / l
}

// AbsoluteDeviation x
func (d Data) AbsoluteDeviation(val float64) float64 {
	return math.Abs(val - d.Mean())
}

// ZScore x
func (d Data) ZScore(val float64) float64 {
	return (val - d.Mean()) / d.StdPop()
}

// PearsonModeSkew x
func (d Data) PearsonModeSkew() (float64, error) {
	if d.IsNormalDist() {
		return 0.0, errors.New("Data is Normally Distributed")
	}
	s := (d.Mean() - d.Mode()) / d.StdPop()
	return s, nil
}

// PearsonSkewCoefficients x
func (d Data) PearsonSkewCoefficients() ([]float64, error) {
	if d.IsNormalDist() {
		return nil, errors.New("Data is Normally Distributed")
	}
	return []float64{
		((3 * d.Mean()) - (3 * d.Mode())) / d.StdPop(),
		((3 * d.Mean()) - (3 * d.Median())) / d.StdPop(),
	}, nil
}
