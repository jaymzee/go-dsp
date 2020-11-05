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

	x, err := wav.ReadFloat64(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}
}
