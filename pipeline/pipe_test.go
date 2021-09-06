package pipeline

import (
	"fmt"
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
