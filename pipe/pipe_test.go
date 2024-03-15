package pipe

import (
	"fmt"
	"math"
)

func ExampleSquare() {
	// set up the pipeline
	c := Gen(1, 2, 3)
	out := Square(c)

	// consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	// Output:
	// 1
	// 4
	// 9
}

func ExampleSinwave() {
	c := Sin(math.Pi/2, 1.0)

	fmt.Printf("%.3f\n", <-c)
	fmt.Printf("%.3f\n", <-c)
	fmt.Printf("%.3f\n", <-c)
	fmt.Printf("%.3f\n", <-c)
	// Output:
	// 0.000
	// 1.000
	// 0.000
	// -1.000
}
