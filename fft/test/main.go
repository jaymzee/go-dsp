package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/fft"
)

func printArray(name string, s []complex128) {
	for i, num := range s {
		fmt.Printf("%s[%d] = %v\n", name, i, num)
	}
}

func main() {
	x := [8]complex128{1, 2, 3, 4, 3, 2, 1, 0}
	x1 := x[:]
	X := fft.FFT(x1)
	x2 := fft.IFFT(X)

	printArray("x", x1)

	fmt.Println("X = fft(x)")
	printArray("X", X)

	fmt.Println("x = fft(X)")
	printArray("x", x2)
}
