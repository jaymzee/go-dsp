package wave

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Read reads a wav file into a Wave struct
func Read(fname string) (*Wave, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wav := new(Wave)
	err = readRIFF(wav, fname, file)
	if err != nil {
		return nil, err
	}
	return wav, nil
}

func readRIFF(wav *Wave, fname string, file *os.File) error {
	chunk := make([]byte, 4)
	err := binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "RIFF" {
		return fmt.Errorf("%s: expected chunk RIFF", fname)
	}
	file.Seek(4, io.SeekCurrent)
	err = binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "WAVE" {
		return fmt.Errorf("%s: expected chunk WAVE", fname)
	}
	for {
		err = binary.Read(file, binary.LittleEndian, &chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s: expected chunk", fname)
		}
		if string(chunk) == "fmt " {
			err = readRIFFfmt(wav, fname, file)
			if err != nil {
				return err
			}
		} else if string(chunk) == "data" {
			var datasize uint32
			err = binary.Read(file, binary.LittleEndian, &datasize)
			if err != nil {
				return fmt.Errorf("%s: expected data size", fname)
			}
			wav.Data = make([]byte, datasize)
			nbytes, err := file.Read(wav.Data)
			if err != nil {
				return err
			}
			if uint32(nbytes) != datasize {
				return fmt.Errorf("%s: data chunk truncated", fname)
			}
		} else {
			fmt.Fprintf(os.Stderr, "%s: ignoring chunk %v\n", fname, chunk)
			var chunksize uint32
			err = binary.Read(file, binary.LittleEndian, &chunksize)
			if err != nil {
				return fmt.Errorf("%s: expected chunk size", fname)
			}
			file.Seek(int64(chunksize), io.SeekCurrent)
		}
	}
	return nil
}

func readRIFFfmt(wav *Wave, fname string, file *os.File) error {
	var fmtSize, bytecount uint32
	err := binary.Read(file, binary.LittleEndian, &fmtSize)
	if err != nil {
		return fmt.Errorf("%s: expected fmt size", fname)
	}
	bytecount += 4
	if fmtSize >= fmtSizeMin {
		err = binary.Read(file, binary.LittleEndian, &wav.Format)
		if err != nil {
			return fmt.Errorf("%s: expected fmt format", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &wav.Channels)
		if err != nil {
			return fmt.Errorf("%s: expected fmt channels", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &wav.SampleRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt samplerate", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &wav.ByteRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt byterate", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &wav.BlockAlign)
		if err != nil {
			return fmt.Errorf("%s: expected fmt blockalign", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &wav.BitsPerSample)
		if err != nil {
			return fmt.Errorf("%s: expected fmt bitspersample", fname)
		}
		bytecount += fmtSizeMin
	} else {
		fmt.Fprintf(os.Stderr,
			"%s: expected length of chunk fmt >= %v bytes, got %v\n",
			fname, fmtSizeMin, fmtSize)
	}
	if bytecount < fmtSize+4 {
		skip := fmtSize + 4 - bytecount
		fmt.Fprintf(os.Stderr,
			"%s: skipping extra %v bytes at end of chunk fmt\n", fname, skip)
		file.Seek(int64(skip), io.SeekCurrent)
		bytecount += skip
	}
	return nil
}
