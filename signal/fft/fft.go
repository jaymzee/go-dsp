// Package fft provides functions for computing the Fast Fourier Transform
package fft

import (
	"math"
	"math/cmplx"
)

// twiddle returns exp(2πj/N)
func twiddle(N int) complex128 {
	angle := 2 * math.Pi / float64(N)
	return cmplx.Exp(complex(0, angle))
}

// reverseBits reverses the order of the bits in x with bit width w.
func reverseBits(x uint32, w int) uint32 {
	x = (x&0xaaaaaaaa)>>1 | (x&0x55555555)<<1
	x = (x&0xcccccccc)>>2 | (x&0x33333333)<<2
	x = (x&0xf0f0f0f0)>>4 | (x&0x0f0f0f0f)<<4
	x = (x&0xff00ff00)>>8 | (x&0x00ff00ff)<<8
	x = x>>16 | x<<16
	return x >> (32 - uint(w))
}

// Shuffle shuffles elements of x by calling Flip on the index of x.
func Shuffle(x []complex128) []complex128 {
	N := len(x)
	w := Log2(N)
	y := make([]complex128, N)
	for n, v := range x {
		y[reverseBits(uint32(n), w)] = v
	}
	return y
}

// IterativeFFT computes the FFT or inverse FFT of x in-place.
// The sign is the sign of the angle of the twiddle factor exp(2πj/N) and
// should be -1 is for FFT and 1 for the inverse FFT.
// The algorithm is based on Data reordering, bit reversal, and in-place
// algorithms section of
// https://en.wikipedia.org/wiki/Cooley-Tukey_FFT_algorithm
//  algorithm iterative-fft is
//     input: Array a of n complex values where n is a power of 2.
//     output: Array A the DFT of a.
//
//     bit-reverse-copy(a, A)
//     n ← a.length
//     for s = 1 to log(n) do
//         m ← 2^s
//         ωm ← exp(−2πi/m)
//         for k = 0 to n-1 by m do
//             ω ← 1
//             for j = 0 to m/2 – 1 do
//                 t ← ω A[k + j + m/2]
//                 u ← A[k + j]
//                 A[k + j] ← u + t
//                 A[k + j + m/2] ← u – t
//                 ω ← ω ωm
//     return A
func IterativeFFT(x []complex128, sign int) {
	N := len(x)
	log2N := Log2(N)
	for s := 1; s <= log2N; s++ {
		m := 1 << uint(s)
		m2 := m >> 1
		Wm := twiddle(sign * m)
		for k := 0; k < N; k += m {
			W := 1 + 0i
			for j := 0; j < m2; j++ {
				t := x[k+j]
				u := W * x[k+j+m2]
				x[k+j] = t + u
				x[k+j+m2] = t - u
				W *= Wm
			}
		}
	}

	// if we are performing an IFFT, then x must be divided by N
	if sign > 0 {
		N := float64(len(x))
		for n, xn := range x {
			x[n] = complex(real(xn)/N, imag(xn)/N)
		}
	}
}

// FFT returns the Fast Fourier Transform of x. Under the hood it uses an
// iterative in place algorigthm with N log N time complexity.
func FFT(x []complex128) []complex128 {
	X := Shuffle(x)
	IterativeFFT(X, -1)
	return X
}

// IFFT returns the Inverse Fast Fourier Transform of X. Under the hood it
// uses an iterative in place algorigthm with N log N time complexity.
func IFFT(X []complex128) []complex128 {
	x := Shuffle(X)
	IterativeFFT(x, 1)
	return x
}
