package main

import (
	"encoding/json"
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/term"
	"os"
	"strconv"
	"strings"
)

// the global config
var cfg Config

// Config is the app configuration based on Flags and Environment variables
type Config struct {
	// command line flags
	PlotFFT     bool
	PlotPCM     bool
	PlotRMS     bool
	PlotLogRMS  float64
	PrettyPrint bool
	PrintFloat  bool
	RangeString string

	// below values are computed in Init() method

	// true if a plot will be rendered
	Plot bool
	// range of samples to print or plot [start:stop]
	Range Range
	// terminal configuration
	Terminal TerminalConfig
}

// Range is the range of a slice [start:stop]
type Range struct {
	Start int
	Stop  int
}

// NewRange returns a new range [start:stop]
func NewRange(start, stop int) Range {
	return Range{Start: start, Stop: stop}
}

// TermConfig is the terminal configuration
type TerminalConfig struct {
	Graphics TerminalGraphics
	Rows     int
	Cols     int
	Xres     int
	Yres     int
}

// TermGraphics is the type of terminal graphics rendering
type TerminalGraphics int

const (
	// ASCIIArt render plots with ASCII art
	ASCIIArt TerminalGraphics = iota
	// Kitty render plots using kitty terminal graphics protocol
	Kitty
	// ITerm2 render plots using iTerm2 graphics protocol (supports mintty)
	ITerm2
	// ConsoleFB render plots directly to the Linux console framebuffer
	ConsoleFB
)

func (tg TerminalGraphics) String() string {
	return [...]string{"ASCIIArt", "Kitty", "ITerm2", "ConsoleFB"}[tg]
}

// ToJSON returns the configuration in pretty printed JSON
func (c *Config) ToJSON() ([]byte, error) {
	jsondata, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsondata, nil
}

// Init probes the environment to finish initialization of the configuration c
func (c *Config) Init(wf *wavio.File) error {
	if c.PlotPCM || c.PlotRMS || c.PlotLogRMS < 0.0 || c.PlotFFT {
		c.Plot = true
	}

	c.Range = parseSampleRange(wf, c.RangeString)

	// detect kitty terminal emulator for terminal graphics
	if strings.Contains(os.Getenv("TERM"), "kitty") && term.Isatty() {
		c.Terminal.Graphics = Kitty
	}

	// detect linux console framebuffer for terminal graphics
	if term.Isaconsole() {
		c.Terminal.Graphics = ConsoleFB
	}

	// terminal window size
	ws := term.GetWinsize()
	c.Terminal.Cols = int(ws.Cols)
	c.Terminal.Rows = int(ws.Rows)
	if ws.Xres <= 0 || ws.Yres <= 0 {
		// no IO_CNTRL TIOCGWINSZ, so create sensible defaults
		c.Terminal.Xres, c.Terminal.Yres = 1024, 768
	} else {
		c.Terminal.Xres, c.Terminal.Yres = int(ws.Xres), int(ws.Yres)
	}

	// allow overriding detected environment
	env := os.Getenv("WAVDUMP")
	for _, expr := range strings.Fields(env) {
		sides := strings.Split(expr, "=")
		key, value := strings.ToLower(sides[0]), strings.ToLower(sides[1])
		switch key {
		case "term":
			switch value {
			case "asciiart", "ascii", "text":
				c.Terminal.Graphics = ASCIIArt
			case "kitty":
				c.Terminal.Graphics = Kitty
			case "iterm2", "iterm":
				c.Terminal.Graphics = ITerm2
			case "consolefb", "console":
				c.Terminal.Graphics = ConsoleFB
			default:
				return fmt.Errorf("bad value for terminal: %s", sides[1])
			}
		case "xres":
			xres, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return fmt.Errorf("bad value for xres: %s", sides[1])
			}
			c.Terminal.Xres = int(xres)
		case "yres":
			yres, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return fmt.Errorf("bad value for yres: %s", sides[1])
			}
			c.Terminal.Yres = int(yres)
		}
	}

	return nil
}

// parse string to get slice start and end values
// values are bounded by the actual number of samples available
func parseSampleRange(wf *wavio.File, str string) Range {
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
	return NewRange(start, end)
}
