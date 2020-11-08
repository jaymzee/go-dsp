package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wav"
	"os"
)

func main() {
	// parse program arguments
	samplesFlag := flag.Int("n", 0, "number of samples to display")
	pcmFlag := flag.Bool("p", false, "display samples as PCM 16")
	shortFlag := flag.Bool("s", false, "display samples on a single line")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] file\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	filename := args[0]

	// read wav file and print the header
	wf, err := wav.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(wf)

	// print some samples from the wav file
	if N := *samplesFlag; N > 0 || *pcmFlag || *shortFlag {
		fmt.Println("data:")
		if *pcmFlag {
			// convert wav file sample data to int16
			x, err := wf.ToInt16(N)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if *shortFlag {
				fmt.Printf("  %#v\n", x)
			} else {
				for n, xn := range x {
					fmt.Printf("  [%d] = %d\n", n, xn)
				}
			}
		} else {
			// convert wav file sample data to floating point
			if wf.BitsPerSample == 32 {
				x, err := wf.ToFloat32(N)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				if *shortFlag {
					fmt.Printf("  %#v\n", x)
				} else {
					for n, xn := range x {
						fmt.Printf("  [%d] = %f\n", n, xn)
					}
				}
			} else {
				// default to 64 bit conversion
				x, err := wf.ToFloat64(N)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				if *shortFlag {
					fmt.Printf("  %#v\n", x)
				} else {
					for n, xn := range x {
						fmt.Printf("  [%d] = %f\n", n, xn)
					}
				}
			}
		}
	}
}
