// Package fft provides functions for computing the Fast Fourier Transform
package fft

// Log2 returns the radix-2 logarithm of integer x using a very fast algorithm
func Log2(x int) int {
	for n := 0; ; n++ {
		x >>= 1
		if x == 0 {
			return n
		}
	}
}

// Conv uses the N point FFT to compute the convolution x and h
func Conv(x, h []complex128, N int) []complex128 {
	// copy to N size array
	xx := make([]complex128, N)
	hh := make([]complex128, N)
	copy(xx, x)
	copy(hh, h)

	// take FFT
	X := Shuffle(xx)
	H := Shuffle(hh)
	IterativeFFT(X, -1)
	IterativeFFT(H, -1)

	// multiply
	for n, Hn := range H {
		X[n] *= Hn
	}

	// IFFT
	y := Shuffle(X)
	IterativeFFT(y, 1)

	return y
}

// Complex converts reals to complex
func Complex(x []float64) []complex128 {
	y := make([]complex128, len(x))
	for n, xn := range x {
		y[n] = complex(xn, 0.0)
	}
	return y
}

// Fmap converts complex to reals by applying f
func Fmap(f func(complex128)float64, x []complex128) []float64 {
	y := make([]float64, len(x))
	for n, xn := range x {
		y[n] = f(xn)
	}
	return y
}
