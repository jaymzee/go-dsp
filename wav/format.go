package wav

import "fmt"

// Format is the wav file encoding format
type Format uint16

const (
	// FormatPCM wav file format is PCM (the default)
	FormatPCM = 1
	// FormatFloat wav file format is IEEE Floating point
	FormatFloat = 3
	// FormatALaw wav file format is A-law companding algorithm
	FormatALaw = 6
	// FormatμLaw wav file format is μ-law companding algorithm
	FormatμLaw = 7
)

func (f Format) String() string {
	var s string
	switch f {
	case FormatPCM:
		s = "PCM"
	case FormatFloat:
		s = "IEEE float"
	case FormatALaw:
		s = "A-law"
	case FormatμLaw:
		s = "μ-law"
	default:
		s = "unknown"
	}
	return fmt.Sprintf("%d (%s)", f, s)
}
