package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
	"os"
)

type Plot struct {
	data []int
	ymin float64
	ymax float64
	W    int
	H    int
	N    int
}

func plotSamples(wf *wavio.File) error {
	winsize, err := GetWinsize()
	if err != nil {
		winsize = &Winsize{24, 80, 0, 0}
	}

	if (useKitty && winsize.Xres > 0 && winsize.Yres > 0) {
		plot, err := WavPlot(wf, int(winsize.Xres), int(winsize.Yres)/4)
		if err != nil {
			return err
		}
		plot.RenderKitty()
	} else {
		plot, err := WavPlot(wf, int(winsize.Cols)-16, int(winsize.Rows)-3)
		if err != nil {
			return err
		}
		plot.RenderASCII(os.Stdout)
	}
	return nil
}

func WavPlot(wf *wavio.File, W int, H int) (*Plot, error) {
	first, last := sampleRange(wf, nFlag)
	N := last - first
	x, err := wf.ToFloat64(first, last)
	if err != nil {
		return nil, err
	}

	// resample to fit screen
	y := make([]float64, W)
	M := (N-1)/W + 1
	i := 0
	j := 0
	t := 0.0
	for _, xn := range x {
		t += xn
		j++
		if j == M {
			y[i] = t / float64(M)
			i++
			j = 0
			t = 0.0
		}
	}
	if j > 0 {
		y[i] = t / float64(j)
	}
	actualW := i

	// rescale and plot
	ymin, ymax := minmax(y[:])
	data := make([]int, W)
	for n, yn := range y {
		data[n] = H - 1 - int((yn-ymin)/(ymax-ymin)*float64(H-1))
	}
	return &Plot{data, ymin, ymax, actualW, H, N}, nil
}

func (plot *Plot) RenderKitty() {
	pixoff := 2
	pixwidth := plot.W + pixoff
	pixbuf := make([]byte, 3 * pixwidth * plot.H)

	for i := 0; i < plot.H; i++ {
		for j := 0; j < min(plot.W, plot.N); j++ {
			if plot.data[j] == i {
				pixbuf[(i*pixwidth + pixoff + j)*3 + 1] = 255;
			}
		}
	}
	writeImage(pixwidth, plot.H, pixbuf)
	fmt.Println()
}

func (plot *Plot) RenderASCII(outf *os.File) {
	for i := 0; i < plot.H; i++ {
		if i == 0 {
			fmt.Fprintf(outf, "\n%11.4e |", plot.ymax)
		} else if i == plot.H-1 {
			fmt.Fprintf(outf, "\n%11.4e |", plot.ymin)
		} else {
			fmt.Fprintf(outf, "\n            |")
		}
		for j := 0; j < min(plot.W, plot.N); j++ {
			if plot.data[j] == i {
				fmt.Fprint(outf, "*")
			} else {
				fmt.Fprint(outf, " ")
			}
		}
	}
	fmt.Fprintln(outf)
}

func minmax(xs []float64) (min float64, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, x := range xs {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return
}
