package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func getBuffer(b []byte, maxBytes int) *bytes.Buffer {
	if maxBytes > 0 && maxBytes < len(b) {
		return bytes.NewBuffer(b[0:maxBytes])
	}
	return bytes.NewBuffer(b)
}

// ToFloat64 converts Data to float64
func (wf *File) ToFloat64(maxSamples int) ([]float64, error) {
	if wf.Channels != 1 {
		return nil, fmt.Errorf("%s: channels must be 1 (mono)", wf.filename)
	}
	if wf.Format == FormatFloat {
		if wf.BitsPerSample == 64 {
			const stride = 8
			buf := getBuffer(wf.Data, maxSamples*stride)
			f64 := make([]float64, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &f64)
			if err != nil {
				return nil, err
			}
			return f64, nil
		}
		if wf.BitsPerSample == 32 {
			const stride = 4
			buf := getBuffer(wf.Data, maxSamples*stride)
			f32 := make([]float32, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &f32)
			if err != nil {
				return nil, err
			}
			f64 := make([]float64, len(f32))
			for n, x := range f32 {
				f64[n] = float64(x)
			}
			return f64, nil
		}
		return nil, fmt.Errorf(
			"%s: IEEE-float must be 32 or 64 bits per sample", wf.filename)
	}
	if wf.Format == FormatPCM {
		if wf.BitsPerSample == 16 {
			const stride = 2
			buf := getBuffer(wf.Data, maxSamples*stride)
			samples := make([]int16, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &samples)
			if err != nil {
				return nil, err
			}
			floats := make([]float64, len(samples))
			for n, v := range samples {
				floats[n] = float64(v) / 32767.0
			}
			return floats, nil
		}
		return nil, fmt.Errorf(
			"%s: PCM must be 16-bit signed", wf.filename)
	}
	return nil, fmt.Errorf("%s: unsupported format %d (%v)",
		wf.filename, wf.Format, wf.Format)
}

// ToFloat32 converts Data to float32
func (wf *File) ToFloat32(maxSamples int) ([]float32, error) {
	if wf.Channels != 1 {
		return nil, fmt.Errorf("%s: channels must be 1 (mono)", wf.filename)
	}
	if wf.Format == FormatFloat {
		if wf.BitsPerSample == 64 {
			const stride = 8
			buf := getBuffer(wf.Data, maxSamples*stride)
			f64 := make([]float64, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &f64)
			if err != nil {
				return nil, err
			}
			f32 := make([]float32, len(f64))
			for n, x := range f64 {
				f32[n] = float32(x)
			}
			return f32, nil
		}
		if wf.BitsPerSample == 32 {
			const stride = 4
			buf := getBuffer(wf.Data, maxSamples*stride)
			floats := make([]float32, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &floats)
			if err != nil {
				return nil, err
			}
			return floats, nil
		}
		return nil, fmt.Errorf(
			"%s: IEEE-float must be 32 or 64 bits per sample", wf.filename)
	}
	if wf.Format == FormatPCM {
		if wf.BitsPerSample == 16 {
			const stride = 2
			buf := getBuffer(wf.Data, maxSamples*stride)
			samples := make([]int16, buf.Len()/stride)
			err := binary.Read(buf, binary.LittleEndian, &samples)
			if err != nil {
				return nil, err
			}
			floats := make([]float32, len(samples))
			for n, v := range samples {
				floats[n] = float32(v) / 32767.0
			}
			return floats, nil
		}
		return nil, fmt.Errorf(
			"%s: PCM must be 16-bit signed", wf.filename)
	}
	return nil, fmt.Errorf("%s: unsupported format %d (%v)",
		wf.filename, wf.Format, wf.Format)
}
