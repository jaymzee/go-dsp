package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
)

func main() {
	// parse program arguments
	samplesFlag := flag.Int("n", 0, "number of samples to print")
	pcmFlag := flag.Bool("p", false, "print samples as PCM 16")
	shortFlag := flag.Bool("s", false, "print samples on a single line")
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
	if *samplesFlag > 0 || *pcmFlag || *shortFlag {
		err := dumpSamples(wf, *samplesFlag, *pcmFlag, *shortFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func dumpSamples(wf *wavio.File, N int, pcm bool, short bool) error {
	const (
		valFmt   = "  %#v\n"
		intFmt   = "  [%d] = %d\n"
		floatFmt = "  [%d] = %f\n"
	)
	fmt.Println("data:")
	if pcm {
		// convert wav file sample data to int16
		x, err := wf.ToInt16(N)
		if err != nil {
			return err
		}
		if short {
			fmt.Printf(valFmt, x)
		} else {
			for n, xn := range x {
				fmt.Printf(intFmt, n, xn)
			}
		}
	} else {
		// convert wav file sample data to floating point
		if wf.BitsPerSample == 32 {
			x, err := wf.ToFloat32(N)
			if err != nil {
				return err
			}
			if short {
				fmt.Printf(valFmt, x)
			} else {
				for n, xn := range x {
					fmt.Printf(floatFmt, n, xn)
				}
			}
		} else { // default to 64 bit conversion
			x, err := wf.ToFloat64(N)
			if err != nil {
				return err
			}
			if short {
				fmt.Printf(valFmt, x)
			} else {
				for n, xn := range x {
					fmt.Printf(floatFmt, n, xn)
				}
			}
		}
	}
	return nil
}
