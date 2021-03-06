package filter

/************************************************************************
* direct form I realization                                             *
*                                                                       *
*  y[n] =   b[0]x[n]   + b[1]x[n-1] + ... + b[L]x[n-L]                  *
*         - a[1]y[n-1] - a[2]y[n-2] - ... - a[M]y[n-M]                  *
*                                                                       *
*              b0                                                       *
* x[n] ----┬---|>-→(+)-------┬---→ y[n]                                 *
*          ↓        ↑        ↓                                          *
*         [z]  b1   |  -a1  [z]                                         *
*          ├---|>-→(+)←-<|---┤                                          *
*          ↓        ↑        ↓                                          *
*         [z]  b2   |  -a2  [z]                                         *
*          └---|>-→(+)←-<|---┘                                          *
*                                                                       *
*************************************************************************/

// DirectForm1 filters xs using the direct form I realization
// b is the numerator quadratic polynomial
// a is the denominator quadratic polynomial
// xs input sample slice
func DirectForm1(b, a, xs []float64) []float64 {
	x := make([]float64, len(b))
	y := make([]float64, len(a))
	ys := make([]float64, len(xs))
	for n, x0 := range xs {
		x[0] = x0
		y0z := 0.0
		for i := len(b) - 1; i >= 0; i-- {
			y0z += b[i] * x[i]
		}
		y0p := 0.0
		for i := len(a) - 1; i > 0; i-- {
			y0p += a[i] * y[i]
		}
		y[0] = y0z - y0p

		for i := len(x) - 1; i > 0; i-- {
			x[i] = x[i-1]
		}
		for i := len(y) - 1; i > 0; i-- {
			y[i] = y[i-1]
		}
		ys[n] = y[0]
	}
	return ys
}
