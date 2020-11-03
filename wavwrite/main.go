package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/wav"
	"math"
)

func main() {
	x := make([]float64, 21)
	x2 := make([]float32, 21)

	for n := 0; n <= 20; n++ {
		x[n] = math.Sin(2.0 * math.Pi * float64(n) / 20.0)
		x2[n] = float32(x[n])
		fmt.Printf("%d %f\n", n, x[n])
	}

	err := wav.WriteFloat64("float64.wav", 8000, x)
	if err != nil {
		panic(err)
	}
	err = wav.WriteFloat32("float32.wav", 8000, x2)
	if err != nil {
		panic(err)
	}
	err = wav.WritePCM16("pcm16.wav", 8000, x)
	if err != nil {
		panic(err)
	}
}
