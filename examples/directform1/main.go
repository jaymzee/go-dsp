package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/filter"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	// process program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s infile outfile\n", os.Args[0])
		os.Exit(1)
	}
	infile := args[0]
	outfile := args[1]

	// read wav file data
	fmt.Println("Reading: ", infile)
	x, fs, err := wavio.ReadFloat64(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// filter it
	b := []float64{0.00425, 0.0, -0.00425}
	a := []float64{1.0, -1.98, 0.991}
	fmt.Printf("Filtering %s with b=%v, a=%v\n", infile, b, a)
	y := filter.DirectForm1(b, a, x)

	// write new wav file
	fmt.Println("Writing: ", outfile)
	err = wavio.Write(outfile, wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
