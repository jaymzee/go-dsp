package wav

import "os"

// ReadFile reads a wav file into memory
func ReadFile(filename string) (*File, error) {
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

// ReadFloat64 reads a wav file and returns the samples as float64
func ReadFloat64(filename string) ([]float64, error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return wf.ToFloat64(0)
}

// ReadFloat32 reads a wav file and returns the samples as float32
func ReadFloat32(filename string) ([]float32, error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return wf.ToFloat32(0)
}
