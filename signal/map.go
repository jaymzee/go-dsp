package signal

// Complex converts reals to complex
func Complex(x []float64) []complex128 {
	y := make([]complex128, len(x))
	for n, xn := range x {
		y[n] = complex(xn, 0.0)
	}
	return y
}

// Map applies f to each real element of x and returns the results
func Map(f func(float64) float64, x []float64) []float64 {
	y := make([]float64, len(x))
	for n, xn := range x {
		y[n] = f(xn)
	}
	return y
}

// MapComplex applies f to each complex element of x and returns the results
func MapComplex(f func(complex128) complex128, x []complex128) []complex128 {
	y := make([]complex128, len(x))
	for n, xn := range x {
		y[n] = f(xn)
	}
	return y
}

// MapReal applies f to each complex element of x and returns the results
func MapReal(f func(complex128) float64, x []complex128) []float64 {
	y := make([]float64, len(x))
	for n, xn := range x {
		y[n] = f(xn)
	}
	return y
}
