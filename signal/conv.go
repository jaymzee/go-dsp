package signal

// Conv returns the convolution of x and h. The computational complexity is
// O(n²) so for large input signals you should use fft.Conv instead.
func Conv(x, h []float64) []float64 {
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
