package main

import (
	"github.com/jaymzee/go-dsp/pipe"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
)

func main() {
	fs := uint32(44100)
	ω := 2.0 * math.Pi * 1000.0 / float64(fs)
	ch := pipe.Sin(ω, 0.9999)
	x := make([]float64, fs)
	for n := 0; n < len(x); n++ {
		x[n] = <-ch
	}
	wavio.Write("sin1k.wav", wavio.PCM, fs, x)
}
