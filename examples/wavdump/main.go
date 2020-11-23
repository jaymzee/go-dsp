package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

var (
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
		fmt.Println(err)
		os.Exit(1)
	}

	// print header
	fmt.Print(wf)

	// print some samples
	if nFlag > 0 || fFlag || lFlag {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func printSamples(wf *wavio.File) error {
	const (
		defaultFmt = "data = %#v\n"
		prettyFmt  = "data = %T{\n"
	)
	N := samples(wf)
	pretty := !lFlag
	if wf.Format == wavio.PCM && !fFlag {
		// convert wav file samples to int16
		x, err := wf.ToInt16(0, N)
		if err != nil {
			return err
		}
		if pretty {
			fmt.Printf(prettyFmt, x)
			for n, xn := range x {
				fmt.Printf("%5d: %6d,\n", n, xn)
			}
			fmt.Println("}")
		} else {
			fmt.Printf(defaultFmt, x)
		}
	} else {
		if wf.BitsPerSample == 64 {
			x, err := wf.ToFloat64(0, N)
			if err != nil {
				return err
			}
			if pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %20.12e,\n", n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		} else {
			// float32 data or PCM data printed as float
			x, err := wf.ToFloat32(0, N)
			if err != nil {
				return err
			}
			if pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %13.6e,\n", n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		}
	}
	return nil
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
