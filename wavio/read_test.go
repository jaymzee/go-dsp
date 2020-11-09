package wavio

import (
	"fmt"
)

func ExampleReadFloat32_f32() {
	x, _, err := ReadFloat32("testdata/float32.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707107
	// [2] = 0.000000
	// [3] = -0.707107
	// [4] = -1.000000
	// [5] = -0.707107
	// [6] = -0.000000
	// [7] = 0.707107
}

func ExampleReadFloat32_f64() {
	x, _, err := ReadFloat32("testdata/float64.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707107
	// [2] = 0.000000
	// [3] = -0.707107
	// [4] = -1.000000
	// [5] = -0.707107
	// [6] = -0.000000
	// [7] = 0.707107
}

func ExampleReadFloat32_pcm16() {
	x, _, err := ReadFloat32("testdata/pcm16.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707083
	// [2] = 0.000000
	// [3] = -0.707083
	// [4] = -1.000000
	// [5] = -0.707083
	// [6] = 0.000000
	// [7] = 0.707083
}

func ExampleReadFloat64_f32() {
	x, _, err := ReadFloat64("testdata/float32.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707107
	// [2] = 0.000000
	// [3] = -0.707107
	// [4] = -1.000000
	// [5] = -0.707107
	// [6] = -0.000000
	// [7] = 0.707107
}

func ExampleReadFloat64_f64() {
	x, _, err := ReadFloat64("testdata/float64.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707107
	// [2] = 0.000000
	// [3] = -0.707107
	// [4] = -1.000000
	// [5] = -0.707107
	// [6] = -0.000000
	// [7] = 0.707107
}

func ExampleReadFloat64_pcm16() {
	x, _, err := ReadFloat64("testdata/pcm16.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %f\n", i, v)
	}

	// Output:
	// [0] = 1.000000
	// [1] = 0.707083
	// [2] = 0.000000
	// [3] = -0.707083
	// [4] = -1.000000
	// [5] = -0.707083
	// [6] = 0.000000
	// [7] = 0.707083
}

func ExampleReadInt16_pcm16() {
	x, _, err := ReadInt16("testdata/pcm16.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %d\n", i, v)
	}

	// Output:
	// [0] = 32767
	// [1] = 23169
	// [2] = 0
	// [3] = -23169
	// [4] = -32767
	// [5] = -23169
	// [6] = 0
	// [7] = 23169
}

func ExampleReadInt16_f32() {
	x, _, err := ReadInt16("testdata/float32.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %d\n", i, v)
	}

	// Output:
	// [0] = 32767
	// [1] = 23169
	// [2] = 0
	// [3] = -23169
	// [4] = -32767
	// [5] = -23169
	// [6] = 0
	// [7] = 23169
}

func ExampleReadInt16_f64() {
	x, _, err := ReadInt16("testdata/float64.wav")
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Printf("[%d] = %d\n", i, v)
	}

	// Output:
	// [0] = 32767
	// [1] = 23169
	// [2] = 0
	// [3] = -23169
	// [4] = -32767
	// [5] = -23169
	// [6] = 0
	// [7] = 23169
}
