package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/term"
	"os"
)

var dumpConfigFlag *bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "environment variables:\n")
		fmt.Fprintf(os.Stderr, "  WAVDUMP=term=iterm xres=800 yres=200    "+
			"terminal graphics (iTerm2 or mintty)\n")
		fmt.Fprintf(os.Stderr,
			"  WAVDUMP=term=text  render plots using ascii art\n")
	}
	flag.BoolVar(&cfg.PrintFloat, "F", false, "print samples as IEEE floats")
	flag.BoolVar(&cfg.PrettyPrint, "P", false, "pretty print samples")
	flag.StringVar(&cfg.RangeString, "N", "",
		"range of samples to print/plot\n"+
			"examples:\n"+
			"  -N 100     first 100 samples\n"+
			"  -N 50:100  50th thru 100th sample\n"+
			"  -N 100:    from 100th sample to the end of the file")
	flag.BoolVar(&cfg.PlotPCM, "p", false, "plot x")
	flag.BoolVar(&cfg.PlotRMS, "r", false, "plot rms(x)")
	flag.BoolVar(&cfg.PlotFFT, "f", false, "plot fft(x) (range must be 2^N)")
	flag.Float64Var(&cfg.PlotLogRMS, "log", 0.0, "plot log(rms(x))\n"+
		"examples:\n  -log=-40   floor >= -40 dB")
	dumpConfigFlag = flag.Bool("config", false, "show configuration")
}

func main() {
	// parse program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	for _, filename := range args {
		dumpFile(filename)
	}
}

func dumpFile(filename string) {
	// read wav file
	wf, err := wavio.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// init configuration (based on flags and the wavfile)
	err = cfg.Init(wf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// for debugging configuration
	if *dumpConfigFlag {
		config, err := cfg.ToJSON()
		if err != nil {
			panic(err)
		}
		fmt.Printf("config = %s\n", config)
	}

	// print summary
	head := fmt.Sprintf("%s: %s [%d:%d]",
		filename, wf.Summary(), cfg.Range.Start, cfg.Range.Stop)
	if !term.Isatty() || len(head) < cfg.Terminal.Cols {
		fmt.Println(head)
	} else {
		fmt.Println(wf.Summary())
	}

	// print some samples
	if (cfg.RangeString != "" && !cfg.Plot) ||
		cfg.PrintFloat || cfg.PrettyPrint {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mdata:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}

	// plot some samples
	if cfg.Plot {
		err := plotWave(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mplot:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}
}
