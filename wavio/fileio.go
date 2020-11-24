package wavio

// Waveform Audio File Format
// RIFF('WAVE'
//      <fmt-ck>           // Format
//      [<fact-ck>]         // Fact chunk
//      [<cue-ck>]          // Cue points
//      [<playlist-ck>]     // Playlist
//      [<assoc-data-list>] // Associated data list
//      <wave-data> )       // Wave data
//
// <wave-data> → data( <bSampleData:Byte> ... )

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Write writes wavio.File struct to a wav file
func (wf *File) Write(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	write := func(v interface{}) {
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, v)
		}
	}

	write([]byte("RIFF"))
	write(uint32(wf.RIFFSize()))
	write([]byte("WAVEfmt "))
	write(uint32(FmtSize))
	write(wf.Format)
	write(wf.Channels)
	write(wf.SampleRate)
	write(wf.ByteRate)
	write(wf.BlockAlign)
	write(wf.BitsPerSample)
	write([]byte("data"))
	write(uint32(len(wf.Data)))
	write(wf.Data)

	return err
}

// read wav file into wavio.File
func (wf *File) readRIFF(file *os.File) error {
	var riffsize uint32
	chunk := make([]byte, 4)
	err := binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "RIFF" {
		return fmt.Errorf("%s: expected RIFF", wf.filename)
	}
	err = binary.Read(file, binary.LittleEndian, &riffsize)
	if err != nil {
		return fmt.Errorf("%s: expected RIFF size", wf.filename)
	}
	err = binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "WAVE" {
		return fmt.Errorf("%s: expected WAVE", wf.filename)
	}
	for {
		err = binary.Read(file, binary.LittleEndian, &chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s: expected chunk", wf.filename)
		}
		if string(chunk) == "fmt " {
			err = wf.readRIFFfmt(file)
			if err != nil {
				return err
			}
		} else if string(chunk) == "data" {
			var datasize uint32
			err = binary.Read(file, binary.LittleEndian, &datasize)
			if err != nil {
				return fmt.Errorf("%s: expected data size", wf.filename)
			}
			wf.Data = make([]byte, datasize)
			nbytes, err := file.Read(wf.Data)
			if err != nil {
				return fmt.Errorf("%s: data expected", wf.filename)
			}
			if uint32(nbytes) != datasize {
				return fmt.Errorf("%s: data truncated", wf.filename)
			}
		} else {
			fmt.Fprintf(os.Stderr, "%s: skipping chunk %s\n",
				wf.filename, chunk)
			var chunksize uint32
			err = binary.Read(file, binary.LittleEndian, &chunksize)
			if err != nil {
				return fmt.Errorf("%s: expected chunk size", wf.filename)
			}
			_, err := file.Seek(int64(chunksize), io.SeekCurrent)
			if err != nil {
				return fmt.Errorf("%s: %s", wf.filename, err)
			}
		}
	}
	return nil
}

// read the wav file fmt chunk into wavio.File
func (wf *File) readRIFFfmt(file *os.File) error {
	var err error
	read := func(expect string, v interface{}) {
		if err == nil && binary.Read(file, binary.LittleEndian, v) != nil {
			err = fmt.Errorf("%s: expected fmt %s", wf.filename, expect)
		}
	}

	var size, bytecount uint32
	read("size", &size)
	bytecount += 4
	if size >= FmtSize {
		read("format", &wf.Format)
		read("channels", &wf.Channels)
		read("samplerate", &wf.SampleRate)
		read("byterate", &wf.ByteRate)
		read("blockalign", &wf.BlockAlign)
		read("bitspersample", &wf.BitsPerSample)
		if err != nil {
			return err
		}
		bytecount += FmtSize
	} else {
		return fmt.Errorf("%s: expected fmt size at least %v bytes, got %v\n",
			wf.filename, FmtSize, size)
	}
	if bytecount < size+4 {
		skip := size + 4 - bytecount
		fmt.Fprintf(os.Stderr, "%s: skipping extra %v bytes at end of fmt\n",
			wf.filename, skip)
		_, err := file.Seek(int64(skip), io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("%s: %s", wf.filename, err)
		}
		bytecount += skip
	}
	return nil
}
