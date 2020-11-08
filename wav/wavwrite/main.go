package main

import (
	"github.com/jaymzee/go-dsp/wav"
	"math"
)

func main() {
	x64 := make([]float64, 21)
	x32 := make([]float32, 21)

	for n := 0; n <= 20; n++ {
		xn := math.Sin(2.0 * math.Pi * float64(n) / 20.0)
		x64[n] = xn
		x32[n] = float32(xn)
	}

	err := wav.WriteFloat64("float64.wav", 8000, x64)
	if err != nil {
		panic(err)
	}
	err = wav.WriteFloat32("float32.wav", 8000, x32)
	if err != nil {
		panic(err)
	}
	err = wav.WritePCM16("pcm16.wav", 8000, x64)
	if err != nil {
		panic(err)
	}
}
