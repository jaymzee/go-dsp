package wav

import (
	"bytes"
	"encoding/binary"
)

// WriteFloat64 writes a float64 slice to a wav file in IEEE float64 format
func WriteFloat64(filename string, rate uint32, data []float64) error {
	const bitsPerSample = 64
	const blockAlign = 1 * bitsPerSample / 8
	wf := File{
		Format:        FormatFloat,
		Channels:      1,
		SampleRate:    rate,
		ByteRate:      rate * 1 * bitsPerSample / 8,
		BlockAlign:    blockAlign,
		BitsPerSample: bitsPerSample,
		Data:          make([]byte, blockAlign*len(data)),
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		return err
	}
	copy(wf.Data, buf.Bytes())
	return wf.Write(filename)
}

// WriteFloat32 writes a float32 slice to a wav file in IEEE float32 format
func WriteFloat32(filename string, rate uint32, data []float32) error {
	const bitsPerSample = 32
	const blockAlign = 1 * bitsPerSample / 8
	wf := File{
		Format:        FormatFloat,
		Channels:      1,
		SampleRate:    rate,
		ByteRate:      rate * 1 * bitsPerSample / 8,
		BlockAlign:    blockAlign,
		BitsPerSample: bitsPerSample,
		Data:          make([]byte, blockAlign*len(data)),
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		return err
	}
	copy(wf.Data, buf.Bytes())
	return wf.Write(filename)
}

// WritePCM16 writes a float64 slice to a wav file in PCM 16-bit format
func WritePCM16(filename string, rate uint32, data []float64) error {
	const bitsPerSample = 16
	const blockAlign = 1 * bitsPerSample / 8
	wf := File{
		Format:        FormatPCM,
		Channels:      1,
		SampleRate:    rate,
		ByteRate:      rate * 1 * bitsPerSample / 8,
		BlockAlign:    blockAlign,
		BitsPerSample: bitsPerSample,
		Data:          make([]byte, blockAlign*len(data)),
	}
	buf := new(bytes.Buffer)
	for _, x := range data {
		y := clamp(x, -1.0, 1.0)
		samp16 := int16(int(32767.0*y+32768.5) - 32768)
		err := binary.Write(buf, binary.LittleEndian, samp16)
		if err != nil {
			return err
		}
	}
	copy(wf.Data, buf.Bytes())
	return wf.Write(filename)
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
