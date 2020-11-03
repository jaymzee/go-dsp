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
	in, err := wave.Read(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(in)
}