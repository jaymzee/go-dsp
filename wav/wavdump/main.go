package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wav"
	"os"
)

func main() {
	// process flags
	samplesFlag := flag.Int("n", 0, "number of samples to display")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-n samples] file\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	filename := args[0]

	// read wav file and print header
	wavfile, err := wav.ReadWave(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(wavfile)

	// print some samples
	if N := *samplesFlag; N > 0 {
		fmt.Println("data:")
		// convert wav file sample data to floating point
		x, err := wavfile.DataFloat64()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for n, xn := range x {
			if n < N {
				fmt.Printf("  [%d] = %f\n", n, xn)
			}
		}
	}
}
