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
	flag.StringVar(&bFlag, "b", "1", "filter numerator")
	flag.StringVar(&aFlag, "a", "1 0.5", "filter denominator")
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
	b, err := parsePoly(bFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "b", err)
		flag.Usage()
		os.Exit(1)
	}
	a, err := parsePoly(aFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "a", err)
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
	y := filter.DirectForm1(b, a, x)

	// write new wav file
	fmt.Println("Writing: ", outfile)
	err = wavio.Write(outfile, wavio.Float, fs, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parsePoly(poly string) (p []float64, err error) {
	split := strings.Fields(poly)
	if len(split) < 1 {
		err = fmt.Errorf("polynomial expected")
		return
	}
	p = make([]float64, len(split))
	for i, str := range split {
		p[i], err = strconv.ParseFloat(str, 64)
		if err != nil {
			err = fmt.Errorf("polynomial malformed")
			return
		}
	}
	return
}
