package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/signal/fft"
	"github.com/jaymzee/go-dsp/wavio"
	"github.com/jaymzee/img/plot"
	"github.com/jaymzee/img/term"
	"github.com/jaymzee/img/term/iTerm2"
	"github.com/jaymzee/img/term/kitty"
	"math"
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

	winsize, err := term.GetWinsize()
	if err != nil {
		winsize = &term.Winsize{24, 80, 0, 0}
	}
	if winsize.Xres <= 0 {
		winsize.Xres = 800
		winsize.Yres = 200
	}

	if cfg.terminal == "kitty" || cfg.terminal == "iTerm" {
		charHeight := winsize.Yres / winsize.Rows
		charWidth := winsize.Xres / winsize.Cols
		width, height = int((winsize.Cols-13)*charWidth), int(charHeight*10)
	} else {
		width, height = int(winsize.Cols)-16, int(winsize.Rows)-3
	}

	x, err := wf.ToFloat64(sampleRange(wf, cfg.nFlag))
	if err != nil {
		return err
	}
	if cfg.fFlag {
		X := fft.Complex(x)
		fft.IterativeFFT(X, 1)
		x = fft.Abs(X)
	}

	if cfg.sFlag < 0 {
		plt = plot.PlotFunc(x, logRms(cfg.sFlag), square, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else if cfg.rFlag {
		plt = plot.PlotFunc(x, math.Sqrt, square, width, height)
		plt.LineColor = 0x0000ffff
		plt.Dots = false
	} else {
		plt = plot.PlotFunc(x, plot.Id, plot.Id, width, height)
		plt.LineColor = 0x00ff00ff
		plt.Dots = true
	}

	if cfg.terminal == "kitty" || cfg.terminal == "iTerm" {
		Plot(plt.RenderPng(), plt.Ymin, plt.Ymax)
	} else {
		fmt.Print(plt.RenderAscii())
	}

	return nil
}

func Plot(buf []byte, min, max float64) {
	fmt.Printf("%11.3e", max)
	if cfg.terminal == "kitty" {
		kitty.WriteImage("a=T,f=100", buf)
	} else {
		iTerm2.WriteImage(buf)
	}
	fmt.Printf("\n\033[A%11.3e\n", min)
}
