# wavio

Package wavio is for reading from and writing to wav audio files

## Waveform Audio File Format

     RIFF('WAVE'
          <fmt-ck>            // Format
          [<fact-ck>]         // Fact chunk
          [<cue-ck>]          // Cue points
          [<playlist-ck>]     // Playlist
          [<assoc-data-list>] // Associated data list
          <wave-data> )       // Wave data

     <wave-data> → data( <bSampleData:Byte> ... )

## Types

### type [File](/file.go#L12)

```go
type File struct {
	Format        Format // format type 1:PCM, 3:FLOAT
	Channels      uint16 // number of channels
	SampleRate    uint32 // sample rate (fs)
	ByteRate      uint32 // byte rate = fs * channels * bitspersample / 8
	BlockAlign    uint16 // block align = channels * bitspersample / 8
	BitsPerSample uint16 // 8, 16, 32 or 64 bits
	Data          []byte // samples
}
```

File contains the raw data for the wav file

### type [Format](/format.go#L6)

```go
type Format uint16
```

Format is the wav file encoding format

## Constants

```go
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

const FmtSize = 16
```
FmtSize is the minimum length of RIFF fmt chunk

## Functions

```go
func NewFile(format Format, channels uint16, sampleRate uint32,
             bitsPerSample uint16, samples int) *File
```
NewFile creates and initializes a new wav file

```go
func ReadFile(filename string) (wf *File, err error)
```
ReadFile reads a wav file into memory

### func [ReadFloat32](/read.go#L39)

```go
func ReadFloat32(filename string) (data []float32, rate uint32, err error)
```
ReadFloat32 reads a wav file and returns the samples as float32

### func [ReadFloat64](/read.go#L25)

```go
func ReadFloat64(filename string) (data []float64, rate uint32, err error)
```
ReadFloat64 reads a wav file and returns the samples as float64

### func [ReadInt16](/read.go#L53)

```go
func ReadInt16(filename string) (data []int16, rate uint32, err error)
```
ReadInt16 reads a wav file and returns the samples as int16

### func [Write](/write.go#L10)

```go
func Write(fname string, format Format, rate uint32, slice interface{}) error
```
Write writes slice to a wav file in the specified format and sample rate


## Methods

```go
func (wf *File) RIFFSize() int
```
RIFFSize computes the RIFF length field of the wav file

```go
func (wf *File) Samples() int
```
Samples computes the total number of samples in each channel

```go
func (wf *File) String() string
```
String formats the header information in the wav file as a string

```go
func (wf *File) Summary() string
```
Summary formats the header information in the wav file as a string but more
terse so it fits on one line

```go
func (wf *File) ToFloat32(start, stop int) (data []float32, err error)
```
ToFloat32 converts Data to float32  
start: start index of slice  
stop: one more than the last index of slice  
if start and stop are both zero, convert the entire slice  
if start is zero and stop > total samples, convert the
entire slice. Otherwise slice the samples before conversion.

```go
func (wf *File) ToFloat64(start, stop int) (data []float64, err error)
```
ToFloat64 converts Data to float64
start: start index of slice  
stop: one more than the last index of slice  
if start and stop are both zero, convert the entire slice  
if start is zero and stop > total samples, convert the
entire slice. Otherwise slice the samples before conversion.

```go
func (wf *File) ToInt16(start, stop int) (data []int16, err error)
ToInt16 converts Data to int16
start: start index of slice  
stop: one more than the last index of slice  
if start and stop are both zero, convert the entire slice  
if start is zero and stop > total samples, convert the
entire slice. Otherwise slice the samples before conversion.
```

```go
func (wf *File) Write(filename string) error
```
Write writes wavio.File struct to a wav file

```go
func (f Format) String() string
```
String converts the Format number to a human readable string

