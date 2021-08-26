package main

import (
	"fmt"
)

func main() {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{4.0, 5.0}

	fmt.Printf("%v\n", convolve(a, b))
	fmt.Printf("%v\n", convolve(b, a))
}

func convolve(x, h []float64) []float64 {
	L := len(x)
	M := len(h)
	y := make([]float64, L+M-1)

	for i := 0; i < len(y); i++ {
		acc := 0.0
		k := min(i, M-1)
		for j := max(0, i-M+1); j < min(L, i+1); j, k = j+1, k-1 {
			acc += h[k] * x[j]
		}
		y[i] = acc
	}

	return y
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
