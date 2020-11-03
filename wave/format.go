package wave

// Format is the wav file encoding format
type Format int

const (
	// FormatPCM wav file format is PCM (the default)
	FormatPCM   Format = 1
	// FormatFloat wav file format is IEEE Floating point
	FormatFloat        = 3
	// FormatALaw wav file format is A-law companding algorithm
	FormatALaw         = 6
	// FormatμLaw wav file format is μ-law companding algorithm
	FormatμLaw         = 7
)

func (f Format) String() string {
	switch f {
	case FormatPCM:
		return "PCM"
	case FormatFloat:
		return "IEEE float"
	case FormatALaw:
		return "A-law"
	case FormatμLaw:
		return "μ-law"
	default:
		return "unknown"
	}
}
