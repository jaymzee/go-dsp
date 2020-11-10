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

	err = binary.Write(file, binary.LittleEndian, []byte("RIFF"))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, uint32(wf.Length()))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, []byte("WAVEfmt "))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, uint32(fmtSizeMin))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.Format)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.Channels)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.SampleRate)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.ByteRate)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.BlockAlign)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.BitsPerSample)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, []byte("data"))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, uint32(len(wf.Data)))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wf.Data)
	if err != nil {
		return err
	}

	return nil
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
				return err
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
				return err
			}
		}
	}
	return nil
}

// read the wav file fmt chunk into wavio.File
func (wf *File) readRIFFfmt(file *os.File) error {
	var fmtSize, bytecount uint32
	err := binary.Read(file, binary.LittleEndian, &fmtSize)
	if err != nil {
		return fmt.Errorf("%s: expected fmt size", wf.filename)
	}
	bytecount += 4
	if fmtSize >= fmtSizeMin {
		err = binary.Read(file, binary.LittleEndian, &wf.Format)
		if err != nil {
			return fmt.Errorf("%s: expected fmt format", wf.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &wf.Channels)
		if err != nil {
			return fmt.Errorf("%s: expected fmt channels", wf.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &wf.SampleRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt samplerate", wf.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &wf.ByteRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt byterate", wf.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &wf.BlockAlign)
		if err != nil {
			return fmt.Errorf("%s: expected fmt blockalign", wf.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &wf.BitsPerSample)
		if err != nil {
			return fmt.Errorf("%s: expected fmt bitspersample", wf.filename)
		}
		bytecount += fmtSizeMin
	} else {
		fmt.Fprintf(os.Stderr, "%s: expected size of fmt >= %v bytes, got %v\n",
			wf.filename, fmtSizeMin, fmtSize)
	}
	if bytecount < fmtSize+4 {
		skip := fmtSize + 4 - bytecount
		fmt.Fprintf(os.Stderr, "%s: skipping extra %v bytes at end of fmt\n",
			wf.filename, skip)
		_, err := file.Seek(int64(skip), io.SeekCurrent)
		if err != nil {
			return err
		}
		bytecount += skip
	}
	return nil
}
