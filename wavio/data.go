package wavio

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	monoConvertError   = "convert to %s: channels must be 1 (mono)"
	floatConvertError  = "convert to %s: IEEE float must be 32 or 64 bits"
	pcmConvertError    = "convert to %s: PCM must be 16-bit signed"
	formatConvertError = "convert to %s: unsupported format %s"
)

const (
	maxInt16 = 32767
	minInt16 = -32768
)

// ToFloat64 converts Data to float64
// start: start index of slice
// stop: one more than the last index of slice
// if start and stop are both zero, convert the entire slice
// if start is zero and stop > total samples, convert the entire slice.
// Otherwise slice the samples before conversion.
func (wf *File) ToFloat64(start, stop int) (data []float64, err error) {
	operation := "float64 " + wf.filename
	if wf.Channels != 1 {
		err = fmt.Errorf(monoConvertError, operation)
		return
	}
	stride := int(wf.BlockAlign)
	buf := getBufferFromSlice(wf.Data, start*stride, stop*stride)
	if wf.Format == Float {
		if wf.BitsPerSample == 64 {
			data = make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		} else if wf.BitsPerSample == 32 {
			x := make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]float64, len(x))
			for n, xn := range x {
				data[n] = float64(xn)
			}
			return
		}
		err = fmt.Errorf(floatConvertError, operation)
		return
	} else if wf.Format == PCM {
		if wf.BitsPerSample == 16 {
			x := make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]float64, len(x))
			for n, xn := range x {
				data[n] = float64(xn) / maxInt16
			}
			return
		}
		err = fmt.Errorf(pcmConvertError, operation)
		return
	}
	err = fmt.Errorf(formatConvertError, operation, wf.Format)
	return
}

// ToFloat32 converts Data to float32
// start: start index of slice
// stop: one more than the last index of slice
// if start and stop are both zero, convert the entire slice
// if start is zero and stop > total samples, convert the entire slice.
// Otherwise slice the samples before conversion.
func (wf *File) ToFloat32(start, stop int) (data []float32, err error) {
	operation := "float 32 " + wf.filename
	if wf.Channels != 1 {
		err = fmt.Errorf(monoConvertError, operation)
		return
	}
	stride := int(wf.BlockAlign)
	buf := getBufferFromSlice(wf.Data, start*stride, stop*stride)
	if wf.Format == Float {
		if wf.BitsPerSample == 64 {
			x := make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]float32, len(x))
			for n, xn := range x {
				data[n] = float32(xn)
			}
			return
		} else if wf.BitsPerSample == 32 {
			data = make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		}
		err = fmt.Errorf(floatConvertError, operation)
		return
	} else if wf.Format == PCM {
		if wf.BitsPerSample == 16 {
			x := make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]float32, len(x))
			for n, xn := range x {
				data[n] = float32(xn) / maxInt16
			}
			return
		}
		err = fmt.Errorf(pcmConvertError, operation)
		return
	}
	err = fmt.Errorf(formatConvertError, operation, wf.Format)
	return
}

// ToInt16 converts Data to int16
// start: start index of slice
// stop: one more than the last index of slice
// if start and stop are both zero, convert the entire slice
// if start is zero and stop > total samples, convert the entire slice.
// Otherwise slice the samples before conversion.
func (wf *File) ToInt16(start, stop int) (data []int16, err error) {
	operation := "int16 " + wf.filename
	if wf.Channels != 1 {
		return nil, fmt.Errorf(monoConvertError, operation)
	}
	stride := int(wf.BlockAlign)
	buf := getBufferFromSlice(wf.Data, start*stride, stop*stride)
	if wf.Format == Float {
		if wf.BitsPerSample == 64 {
			x := make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]int16, len(x))
			for n, xn := range x {
				data[n] = int16(xn * maxInt16)
			}
			return
		}
		if wf.BitsPerSample == 32 {
			x := make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &x)
			if err != nil {
				return
			}
			data = make([]int16, len(x))
			for n, xn := range x {
				data[n] = int16(xn * maxInt16)
			}
			return
		}
		err = fmt.Errorf(floatConvertError, operation)
		return
	}
	if wf.Format == PCM {
		if wf.BitsPerSample == 16 {
			data = make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		}
		err = fmt.Errorf(pcmConvertError, operation)
		return
	}
	err = fmt.Errorf(formatConvertError, operation, wf.Format)
	return
}

// slice b according to the rules described above and return a Buffer
func getBufferFromSlice(b []byte, start, stop int) *bytes.Buffer {
	if start <= 0 && (stop == 0 || stop >= len(b)) {
		return bytes.NewBuffer(b)
	}
	return bytes.NewBuffer(b[start:stop])
}
