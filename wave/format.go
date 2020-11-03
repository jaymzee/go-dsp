package wave

type Format int

const (
	FormatPCM   Format = 1
	FormatFloat        = 3
	FormatALaw         = 6
	FormatμLaw         = 7
)

func (f Format) String() string {
	return [...]string{
		FormatPCM:   "PCM",
		FormatFloat: "IEEE FLOAT",
		FormatALaw:  "ALAW",
		FormatμLaw:  "μLAW",
	}[f]
}
