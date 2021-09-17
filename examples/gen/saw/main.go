package main

import (
	"github.com/jaymzee/go-dsp/pipe"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
)

func main() {
	fs := uint32(44100)
	x := saw(440.0, 0.9999, 20000, fs, fs)
	wavio.Write("sin1k.wav", wavio.PCM, fs, x)
}

func saw(f, r, fb float64, fs uint32, N uint32) []float64 {
	var chans []<-chan float64
	for k:=0; float64(k+1) * f < fb; k++ {
		ω := 2.0 * math.Pi * float64(k+1) * f / float64(fs)
		chans = append(chans, pipe.Sinewave(ω, 0.9999))
	}
	x := make([]float64, N)
	for n := uint32(0); n < N; n++ {
		for _, ch := range chans {
			x[n] += <-ch
		}
		x[n] /= math.Pi
	}
	return x
}
