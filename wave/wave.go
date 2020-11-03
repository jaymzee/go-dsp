package wave

import (
	"fmt"
	"strings"
)

const fmtSizeMin = 16 // minimum length of format block

// Wave contains the raw data for the wav file
type Wave struct {
	Format        uint16 // format type 1:PCM, 3:FLOAT
	Channels      uint16 // number of channels
	SampleRate    uint32 // sample rate (fs)
	ByteRate      uint32 // byte rate = fs * channels * bitspersample / 8
	BlockAlign    uint16 // block align = channels * bitspersample / 8
	BitsPerSample uint16 // 8 or 16 bits
	Data          []byte // data
}

func (wav *Wave) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "format:      %v (%v)\n",
		wav.Format, Format(wav.Format))
	fmt.Fprintf(&b, "channels:    %v\n", wav.Channels)
	fmt.Fprintf(&b, "sample rate: %v\n", wav.SampleRate)
	fmt.Fprintf(&b, "byte rate:   %v\n", wav.ByteRate)
	fmt.Fprintf(&b, "block align: %v\n", wav.BlockAlign)
	fmt.Fprintf(&b, "bits/sample: %v\n", wav.BitsPerSample)
	fmt.Fprintf(&b, "data size:   %v\n", len(wav.Data))
	return b.String()
}

// Length computes RIFF length field of the wav file
func (wav *Wave) Length() int {
	return len(wav.Data) + fmtSizeMin + 2*4 + 3*4 // sizes = 8, tags = 12
}
