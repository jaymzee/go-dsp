package wave

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Wave struct {
	RIFFTag       [4]byte // RIFF chunk tag
	RIFFSize      uint32  // size of file (minus 8 bytes)
	WaveTag       [4]byte // WAVE chunk tag
	FmtTag        [4]byte // 'fmt ' chunk tag
	FmtSize       uint32  // size of data format
	Format        uint16  // format type 1:PCM, 3:FLOAT
	Channels      uint16  // number of channels
	SampleRate    uint32  // sample rate (fs)
	ByteRate      uint32  // byte rate = fs * channels * bitspersample / 8
	BlockAlign    uint16  // block align = channels * bitspersample / 8
	BitsPerSample uint16  // 8 or 16 bits
	DataTag       [4]byte // data chunk tag
	DataSize      uint32  // size of data
	Data          []byte  // data
}

type Format int

const (
	PCM   Format = 1
	Float        = 3
	ALaw         = 6
	MLaw         = 7
)

func (f Format) String() string {
	return [...]string{
		PCM:   "PCM",
		Float: "IEEE FLOAT",
		ALaw:  "ALAW",
		MLaw:  "μLAW",
	}[f]
}

func ReadWaveFile(filename string, wav *Wave) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = binary.Read(file, binary.LittleEndian, &wav.RIFFTag)
	if err != nil || string(wav.RIFFTag[:]) != "RIFF" {
		panic("expected chunk RIFF")
	}
	binary.Read(file, binary.LittleEndian, &wav.RIFFSize)
	err = binary.Read(file, binary.LittleEndian, &wav.WaveTag)
	if err != nil || string(wav.WaveTag[:]) != "WAVE" {
		panic("expected chunk WAVE")
	}
	for {
		var chunktag [4]byte
		err = binary.Read(file, binary.LittleEndian, &chunktag)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("expected chunk")
		}
		if string(chunktag[:]) == "fmt " {
			wav.FmtTag = chunktag
			readFormat(file, wav)
		} else if string(chunktag[:]) == "data" {
			wav.DataTag = chunktag
			err = binary.Read(file, binary.LittleEndian, &wav.DataSize)
			if err != nil {
				panic("expected data size")
			}
			data := make([]byte, wav.DataSize)
			nbytes, err := file.Read(data)
			if err != nil {
				panic(err)
			}
			if uint32(nbytes) != wav.DataSize {
				panic("data chunk truncated")
			}
			wav.Data = data
		} else {
			fmt.Fprintf(os.Stderr, "ignoring chunk %v %s\n",
				chunktag, chunktag)
			var chunksize uint32
			err = binary.Read(file, binary.LittleEndian, &chunksize)
			if err != nil {
				panic("expected chunk size")
			}
			//io.CopyN(ioutil.Discard, file, int64(chunksize))
			file.Seek(int64(chunksize), io.SeekCurrent)
		}
	}
}

func readFormat(file *os.File, wav *Wave) {
	err := binary.Read(file, binary.LittleEndian, &wav.FmtSize)
	var bytecount uint32 = 0
	if err != nil {
		panic("expected fmt size")
	}
	bytecount += 4
	if wav.FmtSize >= 16 {
		binary.Read(file, binary.LittleEndian, &wav.Format)
		binary.Read(file, binary.LittleEndian, &wav.Channels)
		binary.Read(file, binary.LittleEndian, &wav.SampleRate)
		binary.Read(file, binary.LittleEndian, &wav.ByteRate)
		binary.Read(file, binary.LittleEndian, &wav.BlockAlign)
		binary.Read(file, binary.LittleEndian, &wav.BitsPerSample)
		bytecount += 16
	} else {
		fmt.Fprintf(os.Stderr, "expected length of chunk fmt >= 16 bytes")
		fmt.Fprintf(os.Stderr, ", got %v\n", wav.FmtSize)
	}
	if bytecount < wav.FmtSize+4 {
		skip := wav.FmtSize + 4 - bytecount
		fmt.Fprintf(os.Stderr,
			"skipping extra %ld bytes at end of chunk fmt\n", skip)
		//io.CopyN(ioutil.Discard, r, int64(skip))
		file.Seek(int64(skip), io.SeekCurrent)
		bytecount += skip
	}
}
