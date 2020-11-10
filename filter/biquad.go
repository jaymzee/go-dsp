package filter

/************************************************************************
* direct form I realization                                             *
*                                                                       *
*  y[n] =   b[0]x[n]   + b[1]x[n-1] + b[2]x[n-2]                        *
*         - a[1]y[n-1] - a[2]y[n-2] - a[2]y[n-2]                        *
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

// BiQuad filters input samples using the direct form I realization
// b is the numerator polynomial
// a is the denominator polynomial
// x is the input samples
func BiQuad(b, a [3]float64, x []float64) []float64 {
	var x1, x2, y1, y2 float64
	y := make([]float64, len(x))
	for n, x0 := range x {
		y0 := b[0]*x0 + b[1]*x1 + b[2]*x2 - a[1]*y1 - a[2]*y2
		y[n] = y0
		x2 = x1
		x1 = x0
		y2 = y1
		y1 = y0
	}
	return y
}
