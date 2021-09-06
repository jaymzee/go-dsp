# signal

## Functions

### func [Complex](/map.go#L4)

`func Complex(x []float64) []complex128`

Complex converts reals to complex

### func [Conv](/conv.go#L5)

`func Conv(x, h []float64) []float64`

Conv returns the convolution of x and h. The computational complexity is
O(n²) so for large input signals you should use fft.Conv instead.

### func [Log2](/math.go#L44)

`func Log2(x int) int`

Log2 returns the radix-2 logarithm of integer x using a very fast algorithm

### func [Map](/map.go#L13)

`func Map(f func(float64) float64, x []float64) []float64`

Map applies f to each real element of x and returns the results

### func [MapC](/map.go#L22)

`func MapC(f func(complex128) complex128, x []complex128) []complex128`

MapC applies f to each complex element of x and returns the results

### func [MapCtof](/map.go#L31)

`func MapCtof(f func(complex128) float64, x []complex128) []float64`

MapCtof applies f to each complex element of x and returns the results

### func [Max](/math.go#L6)

`func Max(x []float64) float64`

Max finds the maximum value in x

### func [Min](/math.go#L17)

`func Min(x []float64) float64`

Min finds the minimum value in x

## Sub Packages

* [fft](./fft): Package fft provides functions for computing the Fast Fourier Transform

* [filter](./filter): Package filter contains various DSP filtering algorithms
