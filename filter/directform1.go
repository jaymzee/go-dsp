package filter

/************************************************************************
* directform1 is a direct form 1 filter where a and b are any degree    *
* a[0] is assumed to be 1.0                                             *
*                                                                       *
* start with the Z transform of a filter with transfer function H(z):   *
*                                                                       *
*  Y(z) = H(z)·X(z)                                                     *
*  Y(z) = N(z)/D(z)·X(z)                                                *
*  D(z)·Y(z) = N(z)·X(z)                                                *
*                                                                       *
* inverse Z transform and collecting terms to one side gives:           *
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

// DirectForm1 filters x using the direct form I realization
// b is the numerator quadratic polynomial
// a is the denominator quadratic polynomial
// x is the input samples
func DirectForm1(b, a, x []float64) []float64 {
	u := make([]float64, len(b))
	v := make([]float64, len(a))
	y := make([]float64, len(x))
	for n, xn := range x {
		u[0] = xn
		v0 := 0.0
		for i := len(b) - 1; i >= 0; i-- {
			v0 += b[i] * u[i]
		}
		for i := len(a) - 1; i > 0; i-- {
			v0 -= a[i] * v[i]
		}
		v[0] = v0

		for i := len(u) - 1; i > 0; i-- {
			u[i] = u[i-1]
		}
		for i := len(v) - 1; i > 0; i-- {
			v[i] = v[i-1]
		}
		y[n] = v0
	}
	return y
}
