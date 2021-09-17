package main

import (
	"github.com/jaymzee/go-dsp/pipe"
	"github.com/jaymzee/go-dsp/wavio"
)

func main() {
	const fs = 16000
	x := pipe.Saw(220.0, 0.9997, fs/2.0, fs)
	y := pipe.Slice(x, fs)
	wavio.Write("sin1k.wav", wavio.PCM, fs, y)
}
