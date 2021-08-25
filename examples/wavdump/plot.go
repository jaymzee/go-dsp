package main

import (
	"bytes"
	"fmt"
	"github.com/jaymzee/go-dsp/fft"
	"github.com/jaymzee/go-dsp/plot"
	"github.com/jaymzee/go-dsp/wavio"
	"image/png"
	"math"
	"math/cmplx"
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
	gfxPlot := terminal=="kitty" && winsize.Xres > 0 && winsize.Yres > 0
	if gfxPlot {
		charHeight := winsize.Yres / winsize.Rows
		charWidth := winsize.Xres / winsize.Cols
		width, height = int((winsize.Cols-13)*charWidth), int(charHeight*10)
	} else {
		width, height = int(winsize.Cols)-16, int(winsize.Rows)-3
	}

	x, err := wf.ToFloat64(sampleRange(wf, nFlag))
	if err != nil {
		return err
	}
	if fFlag {
		xx := make([]complex128, len(x))
		for n, xn := range x {
			xx[n] = complex(xn, 0)
		}
		fft.IterativeFFT(xx, 1)
		for n, xn := range xx {
			x[n] = cmplx.Abs(xn)
		}
	}

	if sFlag < 0 {
		plt = plot.PlotFunc(x, square, logRms(sFlag), width, height)
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
		GraphicsPlot(plt)
	} else {
		plt.RenderAscii(os.Stdout)
	}

	return nil
}

func GraphicsPlot(plt *plot.Plot) {
	fmt.Printf("%11.3e", plt.Ymax)
	img := plt.RenderImage()
	buf := new(bytes.Buffer)
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	err := enc.Encode(buf, img)
	if err != nil {
		panic(err)
	}
	plot.WriteKitty("a=T,f=100", buf.Bytes())
	fmt.Printf("\n\033[A%11.3e\n", plt.Ymin)
}
