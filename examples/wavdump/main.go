package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"os"
	"strconv"
	"strings"
)

var (
	fFlag    bool
	lFlag    bool
	pFlag    bool
	rFlag    bool
	sFlag    float64
	nFlag    string
	useKitty bool
)

const nFlagHelp = `range of samples to print/plot
examples:
    100		first 100 samples
    50:100	50th thru 100th sample
    100:	from 100th sample to the end of the file`

func init() {
	useKitty = strings.Contains(os.Getenv("TERM"), "kitty") && isatty()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wavfile\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&fFlag, "f", false, "print samples as floating point")
	flag.BoolVar(&lFlag, "l", false,
		"print samples on one line (no pretty print)")
	flag.StringVar(&nFlag, "n", "", nFlagHelp)
	flag.BoolVar(&pFlag, "p", false, "plot samples")
	flag.BoolVar(&rFlag, "r", false, "RMS plot")
	flag.Float64Var(&sFlag, "s", 0.0, "Log RMS plot floor dB (e.g. -40 dB)")
}

func main() {
	// parse program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	filename := args[0]

	// read wav file
	wf, err := wavio.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// print summary
	printSummary(wf)

	// print some samples
	if (nFlag != "" && !pFlag) || fFlag || lFlag {
		err := printSamples(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mdata:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}

	// plot some samples
	if pFlag {
		err := plotWave(wf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;31mplot:\x1b[0m %s\n", err)
			os.Exit(1)
		}
	}
}

func printSummary(wf *wavio.File) {
	summary := wf.Summary()
	first, last := sampleRange(wf, nFlag)
	selection := fmt.Sprintf(" [%d:%d]", first, last)
	fmt.Print(summary)
	if getTermWidth() > len(summary)+len(selection) {
		fmt.Print(selection)
	}
	fmt.Println()
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getTermWidth() int {
	winsize, err := GetWinsize()
	if err != nil {
		return 80
	}
	return int(winsize.Cols)
}
