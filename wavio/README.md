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

`type File struct { ... }`

File contains the raw data for the wav file

### type [Format](/format.go#L6)

`type Format uint16`

Format is the wav file encoding format

## Functions

### func [ReadFloat32](/read.go#L39)

`func ReadFloat32(filename string) (data []float32, rate uint32, err error)`

ReadFloat32 reads a wav file and returns the samples as float32

### func [ReadFloat64](/read.go#L25)

`func ReadFloat64(filename string) (data []float64, rate uint32, err error)`

ReadFloat64 reads a wav file and returns the samples as float64

### func [ReadInt16](/read.go#L53)

`func ReadInt16(filename string) (data []int16, rate uint32, err error)`

ReadInt16 reads a wav file and returns the samples as int16

### func [Write](/write.go#L10)

`func Write(fname string, format Format, rate uint32, slice interface{}) error`

Write writes slice to a wav file in the specified format and sample rate

