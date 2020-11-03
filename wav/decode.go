package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ReadFloat64 reads a float64 wav file and returns the samples
func ReadFloat64(filename string) ([]float64, error) {
	w, err := Read(filename)
	if err != nil {
		return nil, err
	}
	if w.Format != FormatFloat {
		return nil, fmt.Errorf("%s: format IEEE float expected", filename)
	}
	if w.BitsPerSample != 64 {
		return nil, fmt.Errorf("%s: 64 bits per sample expected", filename)
	}
	b := bytes.NewBuffer(w.Data)
	floats := make([]float64, len(w.Data)/int(w.BlockAlign))
	err = binary.Read(b, binary.LittleEndian, &floats)
	if err != nil {
		return nil, err
	}
	return floats, nil
}

// ReadFloat32 reads a float32 wav file and returns the samples
func ReadFloat32(filename string) ([]float32, error) {
	w, err := Read(filename)
	if err != nil {
		return nil, err
	}
	if w.Format != FormatFloat {
		return nil, fmt.Errorf("%s: format IEEE float expected", filename)
	}
	if w.BitsPerSample != 32 {
		return nil, fmt.Errorf("%s: 32 bits per sample expected", filename)
	}
	buf := bytes.NewBuffer(w.Data)
	floats := make([]float32, len(w.Data)/int(w.BlockAlign))
	err = binary.Read(buf, binary.LittleEndian, &floats)
	if err != nil {
		return nil, err
	}
	return floats, nil
}

// ReadPCM16 reads a PCM-16 wav file and returns the samples
func ReadPCM16(filename string) ([]float64, error) {
	w, err := Read(filename)
	if err != nil {
		return nil, err
	}
	if w.Format != FormatPCM {
		return nil, fmt.Errorf("%s: format PCM expected", filename)
	}
	if w.BitsPerSample != 16 {
		return nil, fmt.Errorf("%s: 16 bits per sample expected", filename)
	}
	buf := bytes.NewBuffer(w.Data)
	shorts := make([]int16, len(w.Data)/int(w.BlockAlign))
	err = binary.Read(buf, binary.LittleEndian, &shorts)
	if err != nil {
		return nil, err
	}
	floats := make([]float64, len(shorts))
	for n, v := range shorts {
		floats[n] = float64(v) / 32767.0
	}
	return floats, nil
}
