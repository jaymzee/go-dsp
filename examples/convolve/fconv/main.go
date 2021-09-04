package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/go-dsp/signal/fft"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 4 {
		fmt.Println("Usage: fconv N xwavfile hwavfile ywavfile")
		os.Exit(2)
	}

	N, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	x, fs, err := wavio.ReadFloat64(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	h, _, err := wavio.ReadFloat64(args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	y := fft.Real(fft.Convolve(fft.Complex(x), fft.Complex(h), N))

	err = wavio.Write(args[3], wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
