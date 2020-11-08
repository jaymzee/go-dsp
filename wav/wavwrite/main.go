package main

import (
	"github.com/jaymzee/go-dsp/wav"
	"math"
)

func main() {
	x64 := make([]float64, 21)
	x32 := make([]float32, 21)
	x16 := make([]int16, 21)

	for n := 0; n <= 20; n++ {
		xn := math.Sin(2.0 * math.Pi * float64(n) / 20.0)
		x64[n] = xn
		x32[n] = float32(xn)
	}

	err := wav.Write("float64.wav", wav.FormatFloat, 8000, x64)
	if err != nil {
		panic(err)
	}
	err = wav.Write("float32.wav", wav.FormatFloat, 8000, x32)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-d.wav", wav.FormatPCM, 8000, x64)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-s.wav", wav.FormatPCM, 8000, x32)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-i.wav", wav.FormatPCM, 8000, x16)
	if err != nil {
		panic(err)
	}
}
