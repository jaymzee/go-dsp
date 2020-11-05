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
		fmt.Fprintln(os.Stderr, "Usage: wavdump wavfile")
		os.Exit(1)
	}
	wavfile, err := wav.ReadWave(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(wavfile)
}
