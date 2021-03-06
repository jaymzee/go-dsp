package main

import (
	"flag"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
)

func main() {
	fFlag := flag.Float64("f", 1000, "frequency of sine wave")
	nFlag := flag.Int("n", 8, "number of samples")
	rFlag := flag.Uint("r", 8000, "sample rate")
	flag.Parse()
	f := *fFlag
	N := *nFlag
	fs := uint32(*rFlag)

	Ω := 2 * math.Pi * f
	ω := Ω / float64(fs)
	x64 := make([]float64, N)
	x32 := make([]float32, N)
	x16 := make([]int16, N)

	for n := 0; n < N; n++ {
		xn := math.Cos(ω * float64(n))
		x64[n] = xn
		x32[n] = float32(xn)
		x16[n] = int16(32767 * xn)
	}

	err := wavio.Write("float64.wav", wavio.Float, fs, x64)
	if err != nil {
		panic(err)
	}
	err = wavio.Write("float32.wav", wavio.Float, fs, x32)
	if err != nil {
		panic(err)
	}
	err = wavio.Write("pcm16-d.wav", wavio.PCM, fs, x64)
	if err != nil {
		panic(err)
	}
	err = wavio.Write("pcm16-s.wav", wavio.PCM, fs, x32)
	if err != nil {
		panic(err)
	}
	err = wavio.Write("pcm16-i.wav", wavio.PCM, fs, x16)
	if err != nil {
		panic(err)
	}
}
