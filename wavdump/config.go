package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/term"
	"os"
	"strconv"
	"strings"
)

var cfg Config

// Config is the app configuration based on Flags and Environment variables
type Config struct {
	// flags
	srange  string
	floats  bool
	pretty  bool
	plotpcm bool
	plotrms bool
	plotfft bool
	plotlog float64

	// environment
	termGfx  TermGraphics
	termCols int
	termRows int
	termXres int
	termYres int

	// computed values (based on flags)
	// true if a plot will be rendered
	plot bool
	// range of samples to print or plot [start:stop]
	start int
	stop  int
}

// TermGraphics is the type of terminal graphics rendering
type TermGraphics int

const (
	// ASCIIArt render plots with ASCII art
	ASCIIArt TermGraphics = iota
	// Kitty render plots using kitty terminal graphics protocol
	Kitty
	// ITerm2 render plots using iTerm2 graphics protocol (supports mintty)
	ITerm2
	// ConsoleFB render plots directly to the Linux console framebuffer
	ConsoleFB
)

func (tg TermGraphics) String() string {
	return [...]string{"ASCIIArt", "Kitty", "ITerm2", "ConsoleFB"}[tg]
}

// Init probes the environment to finish initialization of the configuration c
func (c *Config) Init(wf *wavio.File) error {
	if c.plotpcm || c.plotrms || c.plotlog < 0.0 || c.plotfft {
		c.plot = true
	}

	c.start, c.stop = parseSampleRange(wf, c.srange)

	// detect kitty terminal emulator for terminal graphics
	if strings.Contains(os.Getenv("TERM"), "kitty") && term.Isatty() {
		c.termGfx = Kitty
	}

	// detect linux console framebuffer for terminal graphics
	if term.Isaconsole() {
		c.termGfx = ConsoleFB
	}

	// terminal window size
	ws := term.GetWinsize()
	c.termCols = int(ws.Cols)
	c.termRows = int(ws.Rows)
	if ws.Xres <= 0 || ws.Yres <= 0 {
		// no IO_CNTRL TIOCGWINSZ, so create sensible defaults
		c.termXres, c.termYres = 1024, 768
	} else {
		c.termXres, c.termYres = int(ws.Xres), int(ws.Yres)
	}

	// allow overriding detected environment
	env := strings.ToLower(os.Getenv("WAVDUMP"))
	for _, expr := range strings.Fields(env) {
		s := strings.Split(expr, "=")
		switch s[0] {
		case "term":
			switch s[1] {
			case "ITerm2":
			case "iTerm2":
			case "iterm":
				c.termGfx = ITerm2
			case "ASCII":
			case "ASCIIArt":
			case "ascii":
			case "text":
				c.termGfx = ASCIIArt
			}
		case "xres":
			val, err := strconv.ParseUint(s[1], 10, 16)
			if err != nil {
				return fmt.Errorf("bad value for xres: %s", s[1])
			}
			c.termXres = int(val)
		case "yres":
			val, err := strconv.ParseUint(s[1], 10, 16)
			if err != nil {
				return fmt.Errorf("bad value for yres: %s", s[1])
			}
			c.termYres = int(val)
		}
	}

	return nil
}

// parse string to get slice start and end values
// values are bounded by the actual number of samples available
func parseSampleRange(wf *wavio.File, str string) (int, int) {
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
