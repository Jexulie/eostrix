package main

// Comb x
func Comb(totalNumber float64, choice float64, repetition bool) float64 {
	if repetition == true {
		return Factorial(totalNumber+choice-1) / (Factorial(choice) * Factorial(totalNumber-1))
	}

	return Factorial(totalNumber) / (Factorial(choice) * Factorial((totalNumber - choice)))
}
