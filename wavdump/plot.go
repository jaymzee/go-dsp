package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/signal"
	"github.com/jaymzee/go-dsp/signal/fft"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/plot"
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

	if cfg.Terminal.Graphics == ASCIIArt {
		width, height = int(cfg.Terminal.Cols)-16, int(cfg.Terminal.Rows)-5
	} else {
		charH := cfg.Terminal.Yres / cfg.Terminal.Rows
		charW := cfg.Terminal.Xres / cfg.Terminal.Cols
		width, height = int((cfg.Terminal.Cols-13)*charW), int(charH*10)
	}

	// this number was chosen as a maximum because it avoids a seg fault if using console (framebuffer)
	if height > 254 {
		height = 254
	}

	x, err := wf.ToFloat64(cfg.Range.Start, cfg.Range.Stop)
	if err != nil {
		return err
	}
	if cfg.PlotFFT {
		x = signal.MapReal(cmplx.Abs, fft.FFT(signal.Complex(x)))
		x = x[:len(x)/2] // upper half is redundant for real signals
	}

	if cfg.PlotLogRMS < 0 {
		plt = plot.Compose(logRms(cfg.PlotLogRMS), square, x, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else if cfg.PlotRMS {
		plt = plot.Compose(math.Sqrt, square, x, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else {
		plt = plot.Compose(plot.ID, plot.ID, x, width, height)
		plt.LineColor = 0x00ff00ff
		plt.Dots = true
	}

	if cfg.Terminal.Graphics == ASCIIArt {
		fmt.Print(plt.RenderASCII())
	} else {
		err := PlotPNG(plt.RenderPNG(), plt.Ymin, plt.Ymax)
		if err != nil {
			return err
		}
	}

	return nil
}

// PlotPNG writes the PNG image to the terminal using the appropriate
// terminal graphics protocol.
func PlotPNG(buf []byte, min, max float64) error {
	fmt.Printf("%11.3e", max)
	switch cfg.Terminal.Graphics {
	case Kitty:
		err := kitty.WriteImage(os.Stdout, "a=T,f=100", buf)
		if err != nil {
			return err
		}
	case ITerm2:
		err := iterm.WriteImage(os.Stdout, "inline=1", buf)
		if err != nil {
			return err
		}
		fmt.Printf("\033[A")
	case ConsoleFB:
		err := fb.WriteImage("/dev/fb0", buf)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("PlotPNG: no implementation for %v",
			cfg.Terminal.Graphics)
	}
	fmt.Printf("\n\033[A%11.3e\n", min)
	return nil
}
