package combinatronic

// Factorial x
func Factorial(num float64) float64 {
	if num <= 1 {
		return 1.0
	}
	return num * Factorial(num-1.0)
}

// Includes x
func Includes(list []int, num int) bool {
	for _, v := range list {
		if num == v {
			return true
		}
	}
	return false
}

// IsSame x
func IsSame(list1, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}
