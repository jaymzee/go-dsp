# signal

## Functions

### func [Conv](/conv.go#L5)

`func Conv(x, h []float64) []float64`

Conv returns the convolution of x and h. The computational complexity is
O(n²) so for large input signals you should use fft.Conv instead.
