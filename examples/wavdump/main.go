package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	// parse program arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	fFlag := flag.Bool("f", false, "print samples as floating point")
	oFlag := flag.Bool("l", false, "print samples on a line (no pretty print)")
	nFlag := flag.Int("n", 0, "number of samples to print")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// read wav file
	wf, err := wavio.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print header
	fmt.Print(wf.String())

	// print some samples
	if *nFlag > 0 || *fFlag || *oFlag {
		err := dumpSamples(wf, *nFlag, *fFlag, !*oFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func dumpSamples(wf *wavio.File, N int, float, pretty bool) error {
	const (
		defaultFmt = "data: %#v\n"
		prettyFmt  = "data: %T{\n"
	)

	if wf.Format == wavio.PCM && !float {
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
