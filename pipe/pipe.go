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

// Mult Multiplier
func Mult(in <-chan float64, a float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		for x := range in {
			out <- a * x
		}
		close(out)
	}()
	return out
}

// Add Adder
func Add(cs ...<-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		for {
			sum := 0.0
			for _, c := range cs {
				sum += <-c
			}
			out <- sum
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

// Slice is a sink that returns the data in the channel as a length N slice
func Slice(in <-chan float64, N int) []float64 {
	out := make([]float64, N)
	for n := 0; n < N; n++ {
		out[n] = <-in
	}
	return out
}

// Sin is a sine wave generator
func Sin(ω, r float64) <-chan float64 {
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

// Saw is a sawtooth wave generator (band limited)
// f frequency
// r decay rate
// fc cutoff frequency
// fs sample rate
func Saw(f, r, fc, fs float64) <-chan float64 {
	var modes []<-chan float64
	for k := 1; float64(k)*f < fc; k++ {
		ω := 2.0 * math.Pi * float64(k) * f / fs
		a := float64(ipow(-1, k)) / float64(k)
		modes = append(modes, Mult(Sin(ω, r), a))
	}
	return Mult(Add(modes...), 1.0/math.Pi)
}

// return b^n
func ipow(b, n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= b
	}
	return r
}
