package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wave"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: wavdump wavfile")
		os.Exit(1)
	}
	var wav wave.Wave
	wave.ReadWaveFile(args[0], &wav)
	fmt.Printf("%#v\n", wav.RIFFTag)
}
