package wav

import (
	"encoding/binary"
	"os"
)

// Write writes the Wave struct to a file
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
