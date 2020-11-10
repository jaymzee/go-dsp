package wavio

import "fmt"

// Format is the wav file encoding format
type Format uint16

const (
	// PCM indicates wav file format is PCM
	PCM Format = 1
	// Float indicates wav file format is IEEE floating point
	Float Format = 3
	// ALaw indicates wav file format is A-law companding algorithm
	ALaw Format = 6
	// MuLaw indicates wav file format is μ-law companding algorithm
	MuLaw Format = 7
)

// String converts the Format number to a human readable string
func (f Format) String() string {
	var s string
	switch f {
	case PCM:
		s = "PCM"
	case Float:
		s = "IEEE float"
	case ALaw:
		s = "A-law"
	case MuLaw:
		s = "μ-law"
	default:
		s = "unknown"
	}
	return fmt.Sprintf("%d (%s)", f, s)
}
