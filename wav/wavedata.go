package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DataFloat64 returns Wave.Data as float64
func (w *Wave) DataFloat64() ([]float64, error) {
	buf := bytes.NewBuffer(w.Data)
	if w.Channels != 1 {
		return nil, fmt.Errorf("%s: channels must be 1 (mono)", w.filename)
	}
	if w.Format == FormatFloat {
		if w.BitsPerSample == 64 {
			f64 := make([]float64, buf.Len()/8)
			err := binary.Read(buf, binary.LittleEndian, &f64)
			if err != nil {
				return nil, err
			}
			return f64, nil
		}
		if w.BitsPerSample == 32 {
			f32 := make([]float32, buf.Len()/4)
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
			"%s: IEEE-float must be 32 or 64 bits per sample", w.filename)
	}
	if w.Format == FormatPCM {
		if w.BitsPerSample == 16 {
			samples := make([]int16, buf.Len()/2)
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
			"%s: PCM must be 16-bit signed", w.filename)
	}
	return nil, fmt.Errorf("%s: unsupported format %d (%v)",
		w.filename, w.Format, w.Format)
}

// DataFloat32 returns Wave.Data as float32
func (w *Wave) DataFloat32() ([]float32, error) {
	buf := bytes.NewBuffer(w.Data)
	if w.Channels != 1 {
		return nil, fmt.Errorf("%s: channels must be 1 (mono)", w.filename)
	}
	if w.Format == FormatFloat {
		if w.BitsPerSample == 64 {
			f64 := make([]float64, buf.Len()/8)
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
		if w.BitsPerSample == 32 {
			floats := make([]float32, buf.Len()/4)
			err := binary.Read(buf, binary.LittleEndian, &floats)
			if err != nil {
				return nil, err
			}
			return floats, nil
		}
		return nil, fmt.Errorf(
			"%s: IEEE-float must be 32 or 64 bits per sample", w.filename)
	}
	if w.Format == FormatPCM {
		if w.BitsPerSample == 16 {
			samples := make([]int16, buf.Len()/2)
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
			"%s: PCM must be 16-bit signed", w.filename)
	}
	return nil, fmt.Errorf("%s: unsupported format %d (%v)",
		w.filename, w.Format, w.Format)
}
