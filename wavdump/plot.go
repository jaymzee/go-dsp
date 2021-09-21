package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/signal"
	"github.com/jaymzee/go-dsp/signal/fft"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/plot"
	"github.com/jaymzee/img/term"
	"github.com/jaymzee/img/term/fb"
	"github.com/jaymzee/img/term/iterm"
	"github.com/jaymzee/img/term/kitty"
	"math"
	"math/cmplx"
	"os"
)

func square(x float64) float64 {
	return x * x
}

func logRms(floor float64) func(float64) float64 {
	return func(x float64) float64 {
		return math.Max(math.Log10(math.Sqrt(x)), floor)
	}
}

func plotWave(wf *wavio.File) error {
	var plt *plot.Plot
	var width, height int

	winsize := term.GetWinsize()
	if winsize.Xres <= 0 {
		// no IO_CNTRL TIOCGWINSZ, so create sensible defaults
		winsize.Xres = 1024
		winsize.Yres = 768
	}
	if cfg.termXres > 0 {
		// allow overriding in environment
		winsize.Xres = cfg.termXres
		winsize.Yres = cfg.termYres
	}

	if cfg.terminal == "kitty" || cfg.terminal == "iterm" || cfg.terminal == "console" {
		charHeight := winsize.Yres / winsize.Rows
		charWidth := winsize.Xres / winsize.Cols
		width, height = int((winsize.Cols-13)*charWidth), int(charHeight*10)
	} else {
		width, height = int(winsize.Cols)-16, int(winsize.Rows)-5
	}

	// this number was chosen as a maximum because it avoids a seg fault if using console (framebuffer)
	if height > 254 {
		height = 254
	}

	x, err := wf.ToFloat64(sampleRange(wf, cfg.srange))
	if err != nil {
		return err
	}
	if cfg.plotfft {
		x = signal.MapReal(cmplx.Abs, fft.FFT(signal.Complex(x)))
		x = x[:len(x)/2] // upper half is redundant for real signals
	}

	if cfg.plotlog < 0 {
		plt = plot.Compose(logRms(cfg.plotlog), square, x, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else if cfg.plotrms {
		plt = plot.Compose(math.Sqrt, square, x, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else {
		plt = plot.Compose(plot.ID, plot.ID, x, width, height)
		plt.LineColor = 0x00ff00ff
		plt.Dots = true
	}

	if cfg.terminal == "kitty" ||
		cfg.terminal == "iterm" ||
		cfg.terminal == "console" {
		err := PlotPNG(plt.RenderPNG(), plt.Ymin, plt.Ymax)
		if err != nil {
			return err
		}
	} else {
		fmt.Print(plt.RenderASCII())
	}

	return nil
}

// PlotPNG writes the PNG image to the terminal using the appropriate
// terminal graphics protocol.
func PlotPNG(buf []byte, min, max float64) error {
	fmt.Printf("%11.3e", max)
	switch cfg.terminal {
	case "kitty":
		err := kitty.WriteImage(os.Stdout, "a=T,f=100", buf)
		if err != nil {
			return err
		}
	case "iterm":
		err := iterm.WriteImage(os.Stdout, "inline=1", buf)
		if err != nil {
			return err
		}
		fmt.Printf("\033[A")
	case "console":
		err := fb.WriteImage("/dev/fb0", buf)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("%q does not support graphics", cfg.terminal)
	}
	fmt.Printf("\n\033[A%11.3e\n", min)
	return nil
}
