// Package wav is for reading from and writing to wav files
// or WAVE file, an audio file format
package wav

import (
	"fmt"
	"strings"
)

const fmtSizeMin = 16 // minimum length of RIFF fmt chunk

// File contains the raw data for the wav file
type File struct {
	filename      string // filename to provide helpful error messages
	Format        Format // format type 1:PCM, 3:FLOAT
	Channels      uint16 // number of channels
	SampleRate    uint32 // sample rate (fs)
	ByteRate      uint32 // byte rate = fs * channels * bitspersample / 8
	BlockAlign    uint16 // block align = channels * bitspersample / 8
	BitsPerSample uint16 // 8, 16, 32 or 64 bits
	Data          []byte // samples
}

func (wf *File) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "format:      %d (%v)\n",
		wf.Format, wf.Format)
	fmt.Fprintf(&b, "channels:    %v\n", wf.Channels)
	fmt.Fprintf(&b, "sample rate: %v\n", wf.SampleRate)
	fmt.Fprintf(&b, "byte rate:   %v\n", wf.ByteRate)
	fmt.Fprintf(&b, "block align: %v\n", wf.BlockAlign)
	fmt.Fprintf(&b, "bits/sample: %v\n", wf.BitsPerSample)
	fmt.Fprintf(&b, "data size:   %v\n", len(wf.Data))
	return b.String()
}

// Length computes RIFF length field of the wav file
func (wf *File) Length() int {
	return len(wf.Data) + fmtSizeMin + 2*4 + 3*4 // sizes = 8, tags = 12
}
