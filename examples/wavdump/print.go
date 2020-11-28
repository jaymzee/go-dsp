package main

import (
	"fmt"
	"github.com/jaymzee/go-dsp/wavio"
)

func printSamples(wf *wavio.File) error {
	const (
		defaultFmt = "data: %#v\n"
		prettyFmt  = "data: %T{\n"
	)
	N := samples(wf)
	pretty := !lFlag
	if wf.Format == wavio.PCM && !fFlag {
		// convert wav file samples to int16
		x, err := wf.ToInt16(0, N)
		if err != nil {
			return err
		}
		if pretty {
			fmt.Printf(prettyFmt, x)
			for n, xn := range x {
				fmt.Printf("%5d: %6d,\n", n, xn)
			}
			fmt.Println("}")
		} else {
			fmt.Printf(defaultFmt, x)
		}
	} else {
		if wf.BitsPerSample == 64 {
			x, err := wf.ToFloat64(0, N)
			if err != nil {
				return err
			}
			if pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %20.12e,\n", n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		} else {
			// float32 data or PCM data printed as float
			x, err := wf.ToFloat32(0, N)
			if err != nil {
				return err
			}
			if pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %13.6e,\n", n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		}
	}
	return nil
}
