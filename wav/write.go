package wav

import (
	"encoding/binary"
	"os"
)

// Write writes the Wave struct to a file
func (wav *Wave) Write(fname string) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, []byte("RIFF"))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, uint32(wav.Length()))
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
	err = binary.Write(file, binary.LittleEndian, wav.Format)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.Channels)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.SampleRate)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.ByteRate)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.BlockAlign)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.BitsPerSample)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, []byte("data"))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, uint32(len(wav.Data)))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, wav.Data)
	if err != nil {
		return err
	}

	return nil
}
