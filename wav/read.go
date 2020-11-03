package wav

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Read reads a wav file into a Wave struct
func Read(filename string) (*Wave, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	w := new(Wave)
	err = readRIFF(w, filename, file)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func readRIFF(w *Wave, fname string, file *os.File) error {
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
			err = readRIFFfmt(w, fname, file)
			if err != nil {
				return err
			}
		} else if string(chunk) == "data" {
			var datasize uint32
			err = binary.Read(file, binary.LittleEndian, &datasize)
			if err != nil {
				return fmt.Errorf("%s: expected data size", fname)
			}
			w.Data = make([]byte, datasize)
			nbytes, err := file.Read(w.Data)
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

func readRIFFfmt(w *Wave, fname string, file *os.File) error {
	var fmtSize, bytecount uint32
	err := binary.Read(file, binary.LittleEndian, &fmtSize)
	if err != nil {
		return fmt.Errorf("%s: expected fmt size", fname)
	}
	bytecount += 4
	if fmtSize >= fmtSizeMin {
		err = binary.Read(file, binary.LittleEndian, &w.Format)
		if err != nil {
			return fmt.Errorf("%s: expected fmt format", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &w.Channels)
		if err != nil {
			return fmt.Errorf("%s: expected fmt channels", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &w.SampleRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt samplerate", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &w.ByteRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt byterate", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &w.BlockAlign)
		if err != nil {
			return fmt.Errorf("%s: expected fmt blockalign", fname)
		}
		err = binary.Read(file, binary.LittleEndian, &w.BitsPerSample)
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
