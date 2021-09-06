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
		got := Log2(x)
		if got != y {
			t.Errorf("log2(%d) = %d; want %d", x, got, y)
		}
	}
}

func TestReverseBits_width3(t *testing.T) {
	table := [8]uint32{0, 4, 2, 6, 1, 5, 3, 7}
	for x, y := range table {
		got := reverseBits(uint32(x), 3)
		if got != y {
			t.Errorf("reverse(0x%08X) = 0x%08X; want 0x%08X", x, got, y)
		}
	}
}

func TestReverseBits_width32(t *testing.T) {
	table := map[uint32]uint32{
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
	for x, y := range table {
		got := reverseBits(x, 32)
		if got != y {
			t.Errorf("reverse(0x%08X) = 0x%08X; want 0x%08X", x, got, y)
		}
	}
}

func ExampleFFT() {
	x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
	X := FFT(x)
	prettyPrint("X", X)
	// Output:
	// X = []complex128{
	//     0: (16+0i),
	//     1: (-4.82842712474619-4.82842712474619i),
	//     2: (0+0i),
	//     3: (0.8284271247461907-0.8284271247461894i),
	//     4: (0+0i),
	//     5: (0.8284271247461901+0.8284271247461903i),
	//     6: (0+0i),
	//     7: (-4.828427124746191+4.82842712474619i),
	// }
}

func ExampleIFFT() {
	x := []complex128{1, 2, 3, 4, 3, 2, 1, 0}
	X := FFT(x)
	x2 := IFFT(X)
	prettyPrint("x", x2)
	// Output:
	// x = []complex128{
	//     0: (1+5.551115123125783e-17i),
	//     1: (1.9999999999999998+1.5700924586837752e-16i),
	//     2: (3-2.220446049250313e-16i),
	//     3: (4-4.440892098500626e-16i),
	//     4: (3-5.551115123125783e-17i),
	//     5: (2-1.5700924586837752e-16i),
	//     6: (1+2.220446049250313e-16i),
	//     7: (0+4.440892098500626e-16i),
	// }
}

func BenchmarkFFT(b *testing.B) {
	x := make([]complex128, 4096)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FFT(x)
	}
}

func prettyPrint(name string, s []complex128) {
	fmt.Printf("%s = %T{\n", name, s)
	for i, num := range s {
		fmt.Printf("%5d: %#v,\n", i, num)
	}
	fmt.Println("}")
}
