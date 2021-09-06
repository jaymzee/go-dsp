package signal

// Min returns the minimum of a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of a and b
func Max(a, b int) int {
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
