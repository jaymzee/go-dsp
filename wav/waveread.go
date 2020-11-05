package wav

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ReadWave reads a wav file into a Wave struct
func ReadWave(filename string) (*Wave, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	w := &Wave{filename: filename}
	if err := w.readRIFF(file); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *Wave) readRIFF(file *os.File) error {
	chunk := make([]byte, 4)
	err := binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "RIFF" {
		return fmt.Errorf("%s: expected chunk RIFF", w.filename)
	}
	_, err = file.Seek(4, io.SeekCurrent)
	if err != nil {
		return err
	}
	err = binary.Read(file, binary.LittleEndian, &chunk)
	if err != nil || string(chunk) != "WAVE" {
		return fmt.Errorf("%s: expected chunk WAVE", w.filename)
	}
	for {
		err = binary.Read(file, binary.LittleEndian, &chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s: expected chunk", w.filename)
		}
		if string(chunk) == "fmt " {
			err = w.readRIFFfmt(file)
			if err != nil {
				return err
			}
		} else if string(chunk) == "data" {
			var datasize uint32
			err = binary.Read(file, binary.LittleEndian, &datasize)
			if err != nil {
				return fmt.Errorf("%s: expected data size", w.filename)
			}
			w.Data = make([]byte, datasize)
			nbytes, err := file.Read(w.Data)
			if err != nil {
				return err
			}
			if uint32(nbytes) != datasize {
				return fmt.Errorf("%s: data chunk truncated", w.filename)
			}
		} else {
			fmt.Fprintf(os.Stderr, "%s: ignoring chunk %v\n", w.filename, chunk)
			var chunksize uint32
			err = binary.Read(file, binary.LittleEndian, &chunksize)
			if err != nil {
				return fmt.Errorf("%s: expected chunk size", w.filename)
			}
			_, err := file.Seek(int64(chunksize), io.SeekCurrent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Wave) readRIFFfmt(file *os.File) error {
	var fmtSize, bytecount uint32
	err := binary.Read(file, binary.LittleEndian, &fmtSize)
	if err != nil {
		return fmt.Errorf("%s: expected fmt size", w.filename)
	}
	bytecount += 4
	if fmtSize >= fmtSizeMin {
		err = binary.Read(file, binary.LittleEndian, &w.Format)
		if err != nil {
			return fmt.Errorf("%s: expected fmt format", w.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &w.Channels)
		if err != nil {
			return fmt.Errorf("%s: expected fmt channels", w.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &w.SampleRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt samplerate", w.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &w.ByteRate)
		if err != nil {
			return fmt.Errorf("%s: expected fmt byterate", w.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &w.BlockAlign)
		if err != nil {
			return fmt.Errorf("%s: expected fmt blockalign", w.filename)
		}
		err = binary.Read(file, binary.LittleEndian, &w.BitsPerSample)
		if err != nil {
			return fmt.Errorf("%s: expected fmt bitspersample", w.filename)
		}
		bytecount += fmtSizeMin
	} else {
		fmt.Fprintf(os.Stderr,
			"%s: expected length of chunk fmt >= %v bytes, got %v\n",
			w.filename, fmtSizeMin, fmtSize)
	}
	if bytecount < fmtSize+4 {
		skip := fmtSize + 4 - bytecount
		fmt.Fprintf(os.Stderr,
			"%s: skipping extra %v bytes at end of chunk fmt\n",
			w.filename, skip)
		_, err := file.Seek(int64(skip), io.SeekCurrent)
		if err != nil {
			return err
		}
		bytecount += skip
	}
	return nil
}
