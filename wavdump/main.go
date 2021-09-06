package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/term"
	"os"
	"strconv"
	"strings"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "environment variables:\n")
		fmt.Fprintf(os.Stderr,
			"  WAVDUMP=term=iterm xres=800 yres=200    terminal graphics (iTerm2 or mintty)\n")
		fmt.Fprintf(os.Stderr,
			"  WAVDUMP=nogfx    disable graphics (Kitty terminal)\n")
	}
	flag.BoolVar(&cfg.eFlag, "e", false, "print samples as floating point")
	flag.BoolVar(&cfg.fFlag, "f", false, "plot FFT (length must be a power of 2)")
	flag.BoolVar(&cfg.lFlag, "l", false,
		"print samples on one line (no pretty print)")
	flag.StringVar(&cfg.nFlag, "n", "",
		"range of samples to print/plot\n"+
			"examples:\n"+
			"  100     first 100 samples\n"+
			"  50:100  50th thru 100th sample\n"+
			"  100:    from 100th sample to the end of the file")
	flag.BoolVar(&cfg.pFlag, "p", false, "plot samples")
	flag.BoolVar(&cfg.rFlag, "r", false, "plot RMS")
	flag.Float64Var(&cfg.sFlag, "s", 0.0, "plot log RMS, floor in dB (-40 dB)")
}

func main() {
	// parse program arguments
	flag.Parse()
	cfg.ProcessFlags()
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

	// print summary
	first, last := sampleRange(wf, cfg.nFlag)
	head := fmt.Sprintf("%s: %s [%d:%d]", filename, wf.Summary(), first, last)
	if !term.Isatty() || len(head) < getTermWidth() {
		fmt.Println(head)
	} else {
		fmt.Println(wf.Summary())
	}

	// print some samples
	if (cfg.nFlag != "" && !cfg.plot) || cfg.eFlag || cfg.lFlag {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mdata:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}

	// plot some samples
	if cfg.plot {
		err := plotWave(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mplot:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}
}

// parse string to get slice start and end values
// values are bounded by the actual number of samples available
func sampleRange(wf *wavio.File, str string) (int, int) {
	start := 0
	end := 0
	s := strings.Split(str, ":")
	if len(s) > 1 {
		start, _ = strconv.Atoi(s[0])
	}
	if len(s) > 0 {
		end, _ = strconv.Atoi(s[len(s)-1])
	}
	if end > 0 {
		end = min(end, wf.Samples())
	} else {
		end = wf.Samples()
	}
	start = min(start, end)
	return start, end
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getTermWidth() int {
	ws := term.GetWinsize()
	return int(ws.Cols)
}
