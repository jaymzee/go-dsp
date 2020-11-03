package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wav"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: wavread wavfile")
		os.Exit(1)
	}
	filename := args[0]

	w, err := wav.ReadHeader(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if w.Format == wav.FormatPCM && w.BitsPerSample == 16 {
		x, err := wav.ReadPCM16(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for i, v := range x {
			fmt.Printf("[%d] = %f\n", i, v)
		}
	} else if w.Format == wav.FormatFloat && w.BitsPerSample == 32 {
		x, err := wav.ReadFloat32(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for i, v := range x {
			fmt.Printf("[%d] = %f\n", i, v)
		}
	} else if w.Format == wav.FormatFloat && w.BitsPerSample == 64 {
		x, err := wav.ReadFloat64(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for i, v := range x {
			fmt.Printf("[%d] = %f\n", i, v)
		}
	} else {
		fmt.Printf("file format %d (%v) %v-bit not supported",
			w.Format, w.Format, w.BitsPerSample)
	}
}
