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
func (wf *File) ToFloat64(maxSamples int) (data []float64, err error) {
	if wf.Channels != 1 {
		err = fmt.Errorf("%s: channels must be 1 (mono)", wf.filename)
		return
	}
	if wf.Format == FormatFloat {
		if wf.BitsPerSample == 64 {
			const stride = 8
			buf := getBuffer(wf.Data, maxSamples*stride)
			data = make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		}
		if wf.BitsPerSample == 32 {
			const stride = 4
			buf := getBuffer(wf.Data, maxSamples*stride)
			f32 := make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &f32)
			if err != nil {
				return
			}
			data = make([]float64, len(f32))
			for n, x := range f32 {
				data[n] = float64(x)
			}
			return
		}
		err = fmt.Errorf("%s: IEEE float must be 32 or 64 bits per sample",
			wf.filename)
		return
	}
	if wf.Format == FormatPCM {
		if wf.BitsPerSample == 16 {
			const stride = 2
			buf := getBuffer(wf.Data, maxSamples*stride)
			samples := make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &samples)
			if err != nil {
				return
			}
			data = make([]float64, len(samples))
			for n, v := range samples {
				data[n] = float64(v) / 32767.0
			}
			return
		}
		err = fmt.Errorf("%s: PCM must be 16-bit signed", wf.filename)
		return
	}
	err = fmt.Errorf("%s: unsupported format %s", wf.filename, wf.Format)
	return
}

// ToFloat32 converts Data to float32
func (wf *File) ToFloat32(maxSamples int) (data []float32, err error) {
	if wf.Channels != 1 {
		err = fmt.Errorf("%s: channels must be 1 (mono)", wf.filename)
		return
	}
	if wf.Format == FormatFloat {
		if wf.BitsPerSample == 64 {
			const stride = 8
			buf := getBuffer(wf.Data, maxSamples*stride)
			f64 := make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &f64)
			if err != nil {
				return
			}
			data = make([]float32, len(f64))
			for n, x := range f64 {
				data[n] = float32(x)
			}
			return
		}
		if wf.BitsPerSample == 32 {
			const stride = 4
			buf := getBuffer(wf.Data, maxSamples*stride)
			data = make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		}
		err = fmt.Errorf("%s: IEEE float must be 32 or 64 bits per sample",
			wf.filename)
		return
	}
	if wf.Format == FormatPCM {
		if wf.BitsPerSample == 16 {
			const stride = 2
			buf := getBuffer(wf.Data, maxSamples*stride)
			i16 := make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &i16)
			if err != nil {
				return
			}
			data = make([]float32, len(i16))
			for n, v := range i16 {
				data[n] = float32(v) / 32767.0
			}
			return
		}
		err = fmt.Errorf("%s: PCM must be 16-bit signed", wf.filename)
		return
	}
	err = fmt.Errorf("%s: unsupported format %s", wf.filename, wf.Format)
	return
}

// ToInt16 converts Data to int16
func (wf *File) ToInt16(maxSamples int) (data []int16, err error) {
	if wf.Channels != 1 {
		return nil, fmt.Errorf("%s: channels must be 1 (mono)", wf.filename)
	}
	if wf.Format == FormatFloat {
		if wf.BitsPerSample == 64 {
			const stride = 8
			buf := getBuffer(wf.Data, maxSamples*stride)
			f64 := make([]float64, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &f64)
			if err != nil {
				return
			}
			data = make([]int16, len(f64))
			for n, x := range f64 {
				data[n] = int16(x * 32767.0)
			}
			return
		}
		if wf.BitsPerSample == 32 {
			const stride = 4
			buf := getBuffer(wf.Data, maxSamples*stride)
			f32 := make([]float32, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &f32)
			if err != nil {
				return
			}
			data = make([]int16, len(f32))
			for n, x := range f32 {
				data[n] = int16(x * 32767.0)
			}
			return
		}
		err = fmt.Errorf("%s: IEEE-float must be 32 or 64 bits per sample",
			wf.filename)
		return
	}
	if wf.Format == FormatPCM {
		if wf.BitsPerSample == 16 {
			const stride = 2
			buf := getBuffer(wf.Data, maxSamples*stride)
			data = make([]int16, buf.Len()/stride)
			err = binary.Read(buf, binary.LittleEndian, &data)
			return
		}
		err = fmt.Errorf("%s: PCM must be 16-bit signed", wf.filename)
		return
	}
	err = fmt.Errorf("%s: unsupported format %s", wf.filename, wf.Format)
	return
}
