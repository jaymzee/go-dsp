package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	// parse program arguments
	nFlag := flag.Int("n", 0, "number of samples to print")
	oFlag := flag.Bool("o", false, "print samples on one line")
	pFlag := flag.Bool("p", false, "print samples as PCM 16")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] file\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
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
	if *nFlag > 0 || *pFlag || *oFlag {
		err := dumpSamples(wf, *nFlag, *pFlag, !*oFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func dumpSamples(wf *wavio.File, N int, pcm bool, pretty bool) error {
	const (
		valFmt   = "  %#v\n"
		pcmFmt   = "  [%d] = %d\n"
		floatFmt = "  [%d] = %f\n"
	)
	fmt.Println("data:")
	if pcm {
		// convert wav file samples to int16
		x, err := wf.ToInt16(N)
		if err != nil {
			return err
		}
		if pretty {
			for n, xn := range x {
				fmt.Printf(pcmFmt, n, xn)
			}
		} else {
			fmt.Printf(valFmt, x)
		}
	} else {
		if wf.BitsPerSample == 32 {
			x, err := wf.ToFloat32(N)
			if err != nil {
				return err
			}
			if pretty {
				for n, xn := range x {
					fmt.Printf(floatFmt, n, xn)
				}
			} else {
				fmt.Printf(valFmt, x)
			}
		} else {
			// default to converting samples to 64-bit floating point
			x, err := wf.ToFloat64(N)
			if err != nil {
				return err
			}
			if pretty {
				for n, xn := range x {
					fmt.Printf(floatFmt, n, xn)
				}
			} else {
				fmt.Printf(valFmt, x)
			}
		}
	}
	return nil
}
