package pipe

import "math"

// Gen is a simple generator of values
func Gen(nums ...float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func Sinewave(ω, r float64) <-chan float64 {
	a1 := -2 * r * math.Cos(ω)
	a2 := r * r
	b1 := r * math.Sin(ω)
	x0 := 1.0          // delta function
	w1, w2 := 0.0, 0.0 // state variables
	y := make(chan float64)
	go func() {
		for {
			w0 := x0 - a1*w1 - a2*w2
			y <- b1 * w1
			// update delays (shift)
			w2 = w1
			w1 = w0
			x0 = 0.0
		}
	}()
	return y
}

// Square squares the received values and emits them on it's own channel
func Square(in <-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()
	return out
}
