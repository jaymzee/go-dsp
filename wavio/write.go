package wavio

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Write writes slice to a wav file in the specified format and sample rate
func Write(fname string, format Format, rate uint32, slice interface{}) error {
	var wf *File
	buf := new(bytes.Buffer)
	switch x := slice.(type) {
	case []float64:
		if format == Float {
			err := binary.Write(buf, binary.LittleEndian, x)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 64, len(x))
		} else if format == PCM {
			for _, xn := range x {
				xn16 := float64toInt16(xn)
				err := binary.Write(buf, binary.LittleEndian, xn16)
				if err != nil {
					return err
				}
			}
			wf = NewFile(format, 1, rate, 16, len(x))
		}
	case []float32:
		if format == Float {
			err := binary.Write(buf, binary.LittleEndian, x)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 32, len(x))
		} else if format == PCM {
			for _, xn := range x {
				xn16 := float64toInt16(float64(xn))
				err := binary.Write(buf, binary.LittleEndian, xn16)
				if err != nil {
					return err
				}
			}
			wf = NewFile(format, 1, rate, 16, len(x))
		}
	case []int16:
		if format == PCM {
			err := binary.Write(buf, binary.LittleEndian, x)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 16, len(x))
		}
	}
	if wf == nil {
		return fmt.Errorf("Write: conversion from %T to %s is not supported",
			slice, format)
	}
	copy(wf.Data, buf.Bytes())
	return wf.Write(fname)
}

func float64toInt16(x float64) int16 {
	return int16(int(32767*clamp(x, -1, 1)+32768.5) - 32768)
}

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
