package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
	"strings"
)

var (
	fFlag bool
	lFlag bool
	tFlag bool
	nFlag int
	useKitty bool
)

func init() {
	useKitty = strings.Contains(os.Getenv("TERM"), "kitty") && isatty()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&fFlag, "f", false, "print samples as floating point")
	flag.BoolVar(&lFlag, "l", false,
		"print samples on one line (no pretty print)")
	flag.IntVar(&nFlag, "n", 0, "number of samples to print/plot")
	if useKitty {
		flag.BoolVar(&tFlag, "t", false, "plot samples in terminal")
	} else {
		flag.BoolVar(&tFlag, "t", false, "plot samples to stdout")
	}
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
	// fmt.Print(wf)
	fmt.Println(wf.Summary())

	// print some samples
	if (nFlag > 0 && !tFlag) || fFlag || lFlag {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mdata:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}

	// plot some samples
	if tFlag {
		err := plotSamples(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mplot:\x1b[0m %s\n", err)
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
