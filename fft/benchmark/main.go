package main

import (
	"fmt"
	"github.com/jaymzee/go/fft"
)

func main() {
	x := make([]complex128, 4096)
	for i := 0; i < 10000; i++ {
		fft.Fft(x)
	}
}

func printArray(name string, s []complex128) {
	for i, num := range s {
		fmt.Printf("%s[%d] = %v\n", name, i, num)
	}
}

