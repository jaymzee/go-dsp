# fft

Package fft provides functions for computing the Fast Fourier Transform

## Functions

### func [Complex](/util.go#L41)

`func Complex(x []float64) []complex128`

Complex converts reals to complex

### func [Conv](/util.go#L15)

`func Conv(x, h []complex128, N int) []complex128`

Conv uses the N point FFT to compute the convolution x and h

### func [FFT](/fft.go#L72)

`func FFT(x []complex128) []complex128`

FFT returns the Fast Fourier Transform of x. Under the hood it uses an
iterative in place algorigthm with N log N time complexity.

```golang
x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
X := FFT(x)
prettyPrint("X", X)
```

 Output:

```
X = []complex128{
    0: (16+0i),
    1: (-4.82842712474619-4.82842712474619i),
    2: (0+0i),
    3: (0.8284271247461907-0.8284271247461894i),
    4: (0+0i),
    5: (0.8284271247461901+0.8284271247461903i),
    6: (0+0i),
    7: (-4.828427124746191+4.82842712474619i),
}
```

### func [Flip](/fft.go#L16)

`func Flip(x uint32, w int) uint32`

Flip reverses the order of the bits in x with bit width w.

### func [Fmap](/util.go#L50)

`func Fmap(f func(complex128) float64, x []complex128) []float64`

Fmap converts complex to reals by applying f

### func [IFFT](/fft.go#L80)

`func IFFT(X []complex128) []complex128`

IFFT returns the Inverse Fast Fourier Transform of X. Under the hood it
uses an iterative in place algorigthm with N log N time complexity.

```golang
x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
X := FFT(x)
x2 := IFFT(X)
prettyPrint("x", x2)
```

 Output:

```
x = []complex128{
    0: (1+5.551115123125783e-17i),
    1: (1.9999999999999998+1.5700924586837752e-16i),
    2: (3-2.220446049250313e-16i),
    3: (4-4.440892098500626e-16i),
    4: (3-5.551115123125783e-17i),
    5: (2-1.5700924586837752e-16i),
    6: (1+2.220446049250313e-16i),
    7: (0+4.440892098500626e-16i),
}
```

### func [IterativeFFT](/fft.go#L42)

`func IterativeFFT(x []complex128, sign int)`

IterativeFFT computes the FFT or inverse FFT of x in-place.
The sign is the sign of the angle of the twiddle factor exp(2πj/N) and
should be -1 is for FFT and 1 for the inverse FFT.
The algorithm is based on Data reordering, bit reversal, and in-place
algorithms section of
[Cooley-Tukey FFT]([https://en.wikipedia.org/wiki/Cooley-Tukey_FFT_algorithm](https://en.wikipedia.org/wiki/Cooley-Tukey_FFT_algorithm))

### func [Log2](/util.go#L5)

`func Log2(x int) int`

Log2 returns the radix-2 logarithm of integer x using a very fast algorithm

### func [Shuffle](/fft.go#L26)

`func Shuffle(x []complex128) []complex128`

Shuffle shuffles elements of x by calling Flip on the index of x.

### func [Twiddle](/fft.go#L10)

`func Twiddle(N int) complex128`

Twiddle returns exp(2πj/N)
