package pipeline

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
