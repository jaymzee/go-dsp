package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
	"math"
)

func plotSamples(wf *wavio.File) error {
	const (
		W = 64
		H = 21
	)
	N := samples(wf)
	x, err := wf.ToFloat64(0, N)
	if err != nil {
		return err
	}

	// resample to fit screen
	var y [W]float64
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
	var plot [W]int
	for n, yn := range y {
		plot[n] = H - 1 - int((yn-ymin)/(ymax-ymin)*float64(H-1))
	}

	// render plot as text
	for i := 0; i < H; i++ {
		if i == 0 {
			fmt.Printf("\n%11.4e |", ymax)
		} else if i == H-1 {
			fmt.Printf("\n%11.4e |", ymin)
		} else {
			fmt.Printf("\n            |")
		}
		for j := 0; j < min(actualW, N); j++ {
			if plot[j] == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()

	return nil
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
