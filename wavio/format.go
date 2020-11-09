package wavio

import "fmt"

// Format is the wav file encoding format
type Format uint16

const (
	// PCM wav file format is PCM (the default)
	PCM Format = 1
	// Float wav file format is IEEE Floating point
	Float Format = 3
	// ALaw wav file format is A-law companding algorithm
	ALaw Format = 6
	// ULaw wav file format is μ-law companding algorithm
	ULaw Format = 7
)

func (f Format) String() string {
	var s string
	switch f {
	case PCM:
		s = "PCM"
	case Float:
		s = "IEEE float"
	case ALaw:
		s = "A-law"
	case ULaw:
		s = "μ-law"
	default:
		s = "unknown"
	}
	return fmt.Sprintf("%d (%s)", f, s)
}
