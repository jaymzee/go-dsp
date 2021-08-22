package main

import (
	"github.com/jaymzee/go-dsp/plot"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
	"os"
)

func idF64(x float64) float64 {
	return x
}

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

	winsize, err := GetWinsize()
	if err != nil {
		winsize = &Winsize{24, 80, 0, 0}
	}
	gfxPlot := useKitty && winsize.Xres > 0 && winsize.Yres > 0
	if gfxPlot {
		width, height = int(winsize.Xres), int(winsize.Yres)/4
	} else {
		width, height = int(winsize.Cols)-16, int(winsize.Rows)-3
	}

	x, err := wf.ToFloat64(sampleRange(wf, nFlag))
	if err != nil {
		return err
	}

	if sFlag < 0 {
		plt = plot.PlotFunc(x, square, logRms(-40), width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else if rFlag {
		plt = plot.PlotFunc(x, square, math.Sqrt, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else {
		plt = plot.PlotFunc(x, idF64, idF64, width, height)
		plt.LineColor = 0x00ff00ff
		plt.Dots = true
	}

	if gfxPlot {
		plt.RenderKitty()
	} else {
		plt.RenderASCII(os.Stdout)
	}

	return nil
}
