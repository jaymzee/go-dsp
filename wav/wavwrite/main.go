package main

import (
	"flag"
	"github.com/jaymzee/go-dsp/wav"
	"math"
)

func main() {
	freqFlag := flag.Float64("f", 1000, "frequency")
	samplesFlag := flag.Int("n", 8, "number of samples")
	rateFlag := flag.Uint("r", 8000, "sample rate")
	flag.Parse()
	N := *samplesFlag
	fs := uint32(*rateFlag)
	ω := 2 * math.Pi * *freqFlag / float64(fs)
	x64 := make([]float64, N)
	x32 := make([]float32, N)
	x16 := make([]int16, N)

	for n := 0; n < N; n++ {
		xn := math.Cos(ω * float64(n))
		x64[n] = xn
		x32[n] = float32(xn)
		x16[n] = int16(32767 * xn)
	}

	err := wav.Write("float64.wav", wav.FormatFloat, fs, x64)
	if err != nil {
		panic(err)
	}
	err = wav.Write("float32.wav", wav.FormatFloat, fs, x32)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-d.wav", wav.FormatPCM, fs, x64)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-s.wav", wav.FormatPCM, fs, x32)
	if err != nil {
		panic(err)
	}
	err = wav.Write("pcm16-i.wav", wav.FormatPCM, fs, x16)
	if err != nil {
		panic(err)
	}
}
