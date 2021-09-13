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
	first, last := sampleRange(wf, cfg.srange)
	if wf.Format == wavio.PCM && !cfg.floats {
		// convert wav file samples to int16
		x, err := wf.ToInt16(first, last)
		if err != nil {
			return err
		}
		if cfg.pretty {
			fmt.Printf(prettyFmt, x)
			for n, xn := range x {
				fmt.Printf("%5d: %6d,\n", first+n, xn)
			}
			fmt.Println("}")
		} else {
			fmt.Printf(defaultFmt, x)
		}
	} else {
		if wf.BitsPerSample == 64 {
			x, err := wf.ToFloat64(first, last)
			if err != nil {
				return err
			}
			if cfg.pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %20.12e,\n", first+n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		} else {
			// float32 data or PCM data printed as float
			x, err := wf.ToFloat32(first, last)
			if err != nil {
				return err
			}
			if cfg.pretty {
				fmt.Printf(prettyFmt, x)
				for n, xn := range x {
					fmt.Printf("%5d: %13.6e,\n", first+n, xn)
				}
				fmt.Println("}")
			} else {
				fmt.Printf(defaultFmt, x)
			}
		}
	}
	return nil
}
