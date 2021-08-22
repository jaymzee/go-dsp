package plot

import (
	"fmt"
	"math"
	"os"
)

type FuncF64 func(float64) float64

type Plot struct {
	data []int
	ymin float64
	ymax float64
	W    int
	H    int
	N    int
	LineColor uint32
	Dots bool
}

func PlotFunc(x []float64, f, g FuncF64, W, H int) *Plot {
	N := len(x)

	// resample to fit screen
	y := make([]float64, W)
	M := (N-1)/W + 1
	i := 0
	j := 0
	t := 0.0
	for _, xn := range x {
		t += f(xn)
		j++
		if j == M {
			y[i] = g(t / float64(M))
			i++
			j = 0
			t = 0.0
		}
	}
	if j > 0 {
		y[i] = g(t / float64(j))
	}
	actualW := i

	// rescale and plot
	ymin, ymax := minmax(y[:])
	data := make([]int, W)
	for n, yn := range y {
		data[n] = H - 1 - int((yn-ymin)/(ymax-ymin)*float64(H-1))
	}
	return &Plot{data, ymin, ymax, actualW, H, N, 0x00ff00ff, true}
}

func (plt *Plot) RenderKitty() {
	pixoff := 2
	pixwidth := plt.W + pixoff
	pixbuf := make([]byte, 3*pixwidth*plt.H)

	for i := 0; i < plt.H; i++ {
		for j := 0; j < min(plt.W, plt.N); j++ {
			y := plt.data[j]
			if i == y && plt.Dots || i <= y {
				pixbuf[(i*pixwidth+pixoff+j)*3] = byte(plt.LineColor >> 24)
				pixbuf[(i*pixwidth+pixoff+j)*3+1] = byte(plt.LineColor >> 16)
				pixbuf[(i*pixwidth+pixoff+j)*3+2] = byte(plt.LineColor >> 8)
			}
		}
	}
	WriteKitty("a=T,f=24", pixbuf, pixwidth, plt.H)
	fmt.Println()
}

func (plt *Plot) RenderASCII(outf *os.File) {
	for i := 0; i < plt.H; i++ {
		if i == 0 {
			fmt.Fprintf(outf, "\n%11.4e |", plt.ymax)
		} else if i == plt.H-1 {
			fmt.Fprintf(outf, "\n%11.4e |", plt.ymin)
		} else {
			fmt.Fprintf(outf, "\n            |")
		}
		for j := 0; j < min(plt.W, plt.N); j++ {
			if plt.data[j] == i {
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
