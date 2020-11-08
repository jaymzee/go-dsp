package wav

import "os"

// ReadFile reads a wav file into memory
func ReadFile(filename string) (wf *File, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wf = &File{filename: filename}
	err = wf.readRIFF(file)
	if err != nil {
		return nil, err
	}
	return
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

// ReadInt16 reads a wav file and returns the samples as int16
func ReadInt16(filename string) ([]int16, error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return wf.ToInt16(0)
}
