package signal

import "math"

// Max finds the maximum of the args given
func Max(args ...float64) float64 {
	y := -math.MaxFloat64
	for _, arg := range args {
		if arg > y {
			y = arg
		}
	}
	return y
}

// Min finds the minimum of the args given
func Min(args ...float64) float64 {
	y := math.MaxFloat64
	for _, arg := range args {
		if arg < y {
			y = arg
		}
	}
	return y
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
