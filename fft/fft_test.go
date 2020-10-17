package fft

import (
	"fmt"
	"testing"
)

func TestLog2(t *testing.T) {
	table := map[int]int{
		1:  0,
		2:  1,
		4:  2,
		8:  3,
		16: 4,
	}
	for x, y := range table {
		got := log2(x)
		if got != y {
			t.Errorf("log2(%d) = %d; want %d", x, got, y)
		}
	}
}

func TestFlip(t *testing.T) {
	table32 := map[uint32]uint32{
		0x00000001: 0x80000000,
		0x00000002: 0x40000000,
		0x00000004: 0x20000000,
		0x00000008: 0x10000000,
		0x00000010: 0x08000000,
		0x00000200: 0x00400000,
		0x00004000: 0x00020000,
		0x00080000: 0x00001000,
		0x00100000: 0x00000800,
		0x02000000: 0x00000040,
		0x40000000: 0x00000002,
	}
	table3 := [8]uint32{0, 4, 2, 6, 1, 5, 3, 7}
	for x, y := range table32 {
		got := flip(x, 32)
		if got != y {
			t.Errorf("flip(0x%08X) = 0x%08X; want 0x%08X", x, got, y)
		}
	}
	for x, y := range table3 {
		got := flip(uint32(x), 3)
		if got != y {
			t.Errorf("flip(0x%08X) = 0x%08X; want 0x%08X", x, got, y)
		}
	}
}

func ExampleFFT() {
	x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
	X := FFT(x)
	printArray("X", X)
	// Output:
	// X[0] = (16+0i)
	// X[1] = (-4.82842712474619-4.82842712474619i)
	// X[2] = (0+0i)
	// X[3] = (0.8284271247461907-0.8284271247461894i)
	// X[4] = (0+0i)
	// X[5] = (0.8284271247461901+0.8284271247461903i)
	// X[6] = (0+0i)
	// X[7] = (-4.828427124746191+4.82842712474619i)
}

func ExampleIFFT() {
	x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
	X := FFT(x)
	x2 := IFFT(X)
	printArray("x", x2)
	// Output:
	// x[0] = (1+5.551115123125783e-17i)
	// x[1] = (1.9999999999999998+1.5700924586837752e-16i)
	// x[2] = (3-2.220446049250313e-16i)
	// x[3] = (4-4.440892098500626e-16i)
	// x[4] = (3-5.551115123125783e-17i)
	// x[5] = (2-1.5700924586837752e-16i)
	// x[6] = (1+2.220446049250313e-16i)
	// x[7] = (0+4.440892098500626e-16i)
}

func BenchmarkFFT(b *testing.B) {
	x := make([]complex128, 4096)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FFT(x)
	}
}

func printArray(name string, s []complex128) {
	for i, num := range s {
		fmt.Printf("%s[%d] = %v\n", name, i, num)
	}
}
