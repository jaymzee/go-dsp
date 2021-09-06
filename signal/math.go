package signal

import "math"

// Max finds the maximum value in x
func Max(x []float64) float64 {
	m := -math.MaxFloat64
	for _, xn := range x {
		if xn > m {
			m = xn
		}
	}
	return m
}

// Min finds the minimum value in x
func Min(x []float64) float64 {
	m := math.MaxFloat64
	for _, xn := range x {
		if xn < m {
			m = xn
		}
	}
	return m
}

// min returns the minimum of a and b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Log2 returns the radix-2 logarithm of integer x using a very fast algorithm
func Log2(x int) int {
	for n := 0; ; n++ {
		x >>= 1
		if x == 0 {
			return n
		}
	}
}
