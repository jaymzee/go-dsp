package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/signal/filter"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("Usage: fir x.wav h.wav out.wav")
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

	a := []float64{1}

	y := filter.DirectForm1(h, a, x)

	err = wavio.Write(args[2], wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
