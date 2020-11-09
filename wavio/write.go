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
	switch data := slice.(type) {
	case []float64:
		if format == Float {
			err := binary.Write(buf, binary.LittleEndian, data)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 64, len(data))
		} else if format == PCM {
			for _, x := range data {
				y := clamp64(x, -1.0, 1.0)
				samp16 := int16(int(32767.0*y+32768.5) - 32768)
				err := binary.Write(buf, binary.LittleEndian, samp16)
				if err != nil {
					return err
				}
			}
			wf = NewFile(format, 1, rate, 16, len(data))
		}
	case []float32:
		if format == Float {
			err := binary.Write(buf, binary.LittleEndian, data)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 32, len(data))
		} else if format == PCM {
			for _, x := range data {
				y := clamp32(x, -1.0, 1.0)
				samp16 := int16(int(32767.0*y+32768.5) - 32768)
				err := binary.Write(buf, binary.LittleEndian, samp16)
				if err != nil {
					return err
				}
			}
			wf = NewFile(format, 1, rate, 16, len(data))
		}
	case []int16:
		if format == PCM {
			err := binary.Write(buf, binary.LittleEndian, data)
			if err != nil {
				return err
			}
			wf = NewFile(format, 1, rate, 16, len(data))
		}
	}
	if wf == nil {
		return fmt.Errorf("Write: conversion from %T to %s is not supported",
			slice, format)
	}
	copy(wf.Data, buf.Bytes())
	return wf.Write(fname)
}

func clamp64(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func clamp32(x float32, min float32, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
