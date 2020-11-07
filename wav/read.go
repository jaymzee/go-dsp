package wav

import "os"

// ReadFloat64 reads the wav file data and returns it as float64
func ReadFloat64(filename string) ([]float64, error) {
	wf, err := Read(filename)
	if err != nil {
		return nil, err
	}
	return wf.ToFloat64(0)
}

// ReadFloat32 reads the wav file data and returns it as float32
func ReadFloat32(filename string) ([]float32, error) {
	wf, err := Read(filename)
	if err != nil {
		return nil, err
	}
	return wf.ToFloat32(0)
}

// Read reads a wav file into a wav.File struct
func Read(filename string) (*File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wf := &File{filename: filename}
	if err := wf.readRIFF(file); err != nil {
		return nil, err
	}
	return wf, nil
}
