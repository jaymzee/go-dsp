package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

var (
	aFlag bool
	fFlag bool
	lFlag bool
	nFlag int
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&aFlag, "a", false, "plot samples as ascii art")
	flag.BoolVar(&fFlag, "f", false, "print samples as floating point")
	flag.BoolVar(&lFlag, "l", false,
		"print samples on one line (no pretty print)")
	flag.IntVar(&nFlag, "n", 0, "number of samples to print")
}

func main() {
	// parse program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	filename := args[0]

	// read wav file
	wf, err := wavio.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// print header
	fmt.Print(wf)

	// print some samples
	if (nFlag > 0 && !aFlag) || fFlag || lFlag {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// plot some samples
	if aFlag {
		err := plotSamples(wf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

// samples is -n flag bounded by the actual number of samples available
func samples(wf *wavio.File) int {
	var count int
	if nFlag > 0 {
		count = min(nFlag, wf.Samples())
	} else {
		// default to selecting all of the samples
		count = wf.Samples()
	}
	return count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
