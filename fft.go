package fft

import (
	"math"
	"math/cmplx"
)

func log2(x int) int {
	for n := 0; ; n++ {
		x >>= 1
		if x == 0 {
			return n
		}
	}
}

func twiddle(N int) complex128 {
	angle := 2 * math.Pi / float64(N)
	return cmplx.Exp(complex(0, angle))
}

func flip(x uint32, w int) uint32 {
	x = (x&0xaaaaaaaa)>>1 | (x&0x55555555)<<1
	x = (x&0xcccccccc)>>2 | (x&0x33333333)<<2
	x = (x&0xf0f0f0f0)>>4 | (x&0x0f0f0f0f)<<4
	x = (x&0xff00ff00)>>8 | (x&0x00ff00ff)<<8
	x = x>>16 | x<<16
	return x >> (32 - uint(w))
}

func shuffle(x []complex128) []complex128 {
	N := len(x)
	w := log2(N)
	y := make([]complex128, N)
	for n, v := range x {
		y[flip(uint32(n), w)] = v
	}
	return y
}

func fftInPlace(x []complex128, sign int) {
	N := len(x)
	log2N := log2(N)
	for s := 1; s <= log2N; s++ {
		m := 1 << uint(s)
		m2 := m >> 1
		Wm := twiddle(sign * m)
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
}

// Fft computes the Fast Fourier Transform
func Fft(x []complex128) []complex128 {
	X := shuffle(x)
	fftInPlace(X, -1)
	return X
}

// Ifft computes the Inverse Fast Fourier Transform
func Ifft(X []complex128) []complex128 {
	x := shuffle(X)
	fftInPlace(x, 1)
	N := float64(len(x))
	for n, v := range x {
		x[n] = complex(real(v)/N, imag(v)/N)
	}
	return x
}
