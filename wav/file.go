package wav

import (
	"fmt"
	"strings"
)

const fmtSizeMin = 16 // minimum length of RIFF fmt chunk

// File contains the raw data for the wav file
type File struct {
	Format        Format // format type 1:PCM, 3:FLOAT
	Channels      uint16 // number of channels
	SampleRate    uint32 // sample rate (fs)
	ByteRate      uint32 // byte rate = fs * channels * bitspersample / 8
	BlockAlign    uint16 // block align = channels * bitspersample / 8
	BitsPerSample uint16 // 8, 16, 32 or 64 bits
	Data          []byte // samples
	filename      string // filename only to provide helpful error messages
}

// NewFile creates and initializes a new wav file
func NewFile(format Format, channels uint16, sampleRate uint32,
	bitsPerSample uint16, samples int) *File {
	blockAlign := channels * bitsPerSample / 8
	return &File{
		Format:        format,
		Channels:      channels,
		SampleRate:    sampleRate,
		ByteRate:      sampleRate * uint32(blockAlign),
		BitsPerSample: bitsPerSample,
		BlockAlign:    blockAlign,
		Data:          make([]byte, int(blockAlign)*samples),
	}
}

// String formats the header information in the wav file as a string
func (wf *File) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "format:      %s\n", wf.Format)
	fmt.Fprintf(&b, "channels:    %d\n", wf.Channels)
	fmt.Fprintf(&b, "sample rate: %d\n", wf.SampleRate)
	fmt.Fprintf(&b, "byte rate:   %d\n", wf.ByteRate)
	fmt.Fprintf(&b, "block align: %d\n", wf.BlockAlign)
	fmt.Fprintf(&b, "bits/sample: %d\n", wf.BitsPerSample)
	fmt.Fprintf(&b, "data size:   %d\n", len(wf.Data))
	return b.String()
}

// Length computes the RIFF length field of the wav file
func (wf *File) Length() int {
	return len(wf.Data) + fmtSizeMin + 2*4 + 3*4 // sizes = 8, tags = 12
}
