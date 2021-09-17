package pipeline

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

func Sinewave(w, r float64) <-chan float64 {
	a := [3]float64{1.0, -2 * r * math.Cos(w), r * r}
	b := r * math.Sin(w)
	out := make(chan float64)
	var v [3]float64
	x0 := 1.0
	go func() {
		for {
			v0 := x0 - a[1]*v[1] - a[2]*v[2]
			out <- b * v[1]
			v[2] = v[1]
			v[1] = v0
			x0 = 0.0
		}
	}()
	return out
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
