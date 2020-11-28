package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/filter"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
	"strconv"
	"strings"
)

var (
	bFlag string
	aFlag string
)

func init() {
	flag.Usage = func() {
		in := "in.wav"
		out := "out.wav"
		name := os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %s [options] infile outfile\n", name)
		fmt.Fprintln(os.Stderr, "options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "examples:")
		fmt.Fprintf(os.Stderr, "  %s -b=%q -a=%q %s %s\n",
			name, "1", "1 0.5", in, out)
		fmt.Fprintf(os.Stderr, "  %s -b=%q -a=%q %s %s\n",
			name, "0.00425 0.0 -0.00425", "1.0 -1.98 0.991", in, out)
	}
	flag.StringVar(&bFlag, "b", "1", "filter numerator (quadratic)")
	flag.StringVar(&aFlag, "a", "1 0.5", "filter denominator (quadratic)")
}

func main() {
	// process program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	infile := args[0]
	outfile := args[1]

	// parse polynomials
	var a, b [3]float64
	if err := parsePoly3(aFlag, &a); err != nil {
		fmt.Fprintln(os.Stderr, "a", err)
		flag.Usage()
		os.Exit(1)
	}
	if err := parsePoly3(bFlag, &b); err != nil {
		fmt.Fprintln(os.Stderr, "b", err)
		flag.Usage()
		os.Exit(1)
	}

	// read wav file data
	fmt.Println("Reading: ", infile)
	x, fs, err := wavio.ReadFloat64(infile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// filter it
	fmt.Printf("Filtering %s with b=%v, a=%v\n", infile, b, a)
	y := filter.BiQuad(b, a, x)

	// write new wav file
	fmt.Println("Writing: ", outfile)
	err = wavio.Write(outfile, wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parsePoly3(poly string, arr *[3]float64) error {
	split := strings.Fields(poly)
	if len(split) == 0 {
		return fmt.Errorf("polynomial expected")
	} else if len(split) > 3 {
		return fmt.Errorf("polynomial has too many coefficents")
	}
	for i, str := range split {
		if c, err := strconv.ParseFloat(str, 64); err == nil {
			arr[i] = c
		} else {
			return fmt.Errorf("polynomial malformed")
		}
	}
	return nil
}
