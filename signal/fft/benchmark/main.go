package main

import "github.com/jaymzee/go-dsp/fft"

func main() {
	x := make([]complex128, 4096)
	for i := 0; i < 10000; i++ {
		fft.FFT(x)
	}
}
