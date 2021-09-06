package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/signal"
	"github.com/jaymzee/go-dsp/signal/fft"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("Usage: fftconv x.wav h.wav out.wav")
		os.Exit(2)
	}

	x, fs, err := wavio.ReadFloat64(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	h, _, err := wavio.ReadFloat64(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	N := 2 << uint(signal.Log2(len(x)+len(h)-2))
	y := signal.MapCtof(Real, fft.Conv(signal.Complex(x), signal.Complex(h), N))

	err = wavio.Write(args[2], wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Real returns the real part of x
func Real(x complex128) float64 {
	return real(x)
}
