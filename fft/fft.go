package fft

import (
	"math"
	"math/cmplx"
)

// Log2 returns the radix-2 logarithm of integer x
// very fast implementation
func Log2(x int) int {
	for n := 0; ; n++ {
		x >>= 1
		if x == 0 {
			return n
		}
	}
}

// Twiddle returns exp(2πj/N)
func Twiddle(N int) complex128 {
	angle := 2 * math.Pi / float64(N)
	return cmplx.Exp(complex(0, angle))
}

// Flip reverses the order of bits in x
//  x value to reverse
//  w width (number of bits)
func Flip(x uint32, w int) uint32 {
	x = (x&0xaaaaaaaa)>>1 | (x&0x55555555)<<1
	x = (x&0xcccccccc)>>2 | (x&0x33333333)<<2
	x = (x&0xf0f0f0f0)>>4 | (x&0x0f0f0f0f)<<4
	x = (x&0xff00ff00)>>8 | (x&0x00ff00ff)<<8
	x = x>>16 | x<<16
	return x >> (32 - uint(w))
}

// Shuffle returns x shuffled
// entries are shuffled by calling flip on the index
func Shuffle(x []complex128) []complex128 {
	N := len(x)
	w := Log2(N)
	y := make([]complex128, N)
	for n, v := range x {
		y[Flip(uint32(n), w)] = v
	}
	return y
}

// IterativeFFT performs in-place and iterative FFT algorithm
func IterativeFFT(x []complex128, sign int) {
	N := len(x)
	log2N := Log2(N)
	for s := 1; s <= log2N; s++ {
		m := 1 << uint(s)
		m2 := m >> 1
		Wm := Twiddle(sign * m)
		for k := 0; k < N; k += m {
			W := complex(1, 0)
			for j := 0; j < m2; j++ {
				t := x[k+j]
				u := W * x[k+j+m2]
				x[k+j] = t + u
				x[k+j+m2] = t - u
				W *= Wm
			}
		}
	}
	// check if we are performing an IFFT
	if sign > 0 {
		N := float64(len(x))
		for n, xn := range x {
			x[n] = complex(real(xn)/N, imag(xn)/N)
		}
	}
}

// FFT computes the Fast Fourier Transform
func FFT(x []complex128) []complex128 {
	X := Shuffle(x)
	IterativeFFT(X, -1)
	return X
}

// IFFT computes the Inverse Fast Fourier Transform
func IFFT(X []complex128) []complex128 {
	x := Shuffle(X)
	IterativeFFT(x, 1)
	return x
}
