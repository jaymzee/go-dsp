package wavio

import "os"

// ReadFile reads a wav file into memory
func ReadFile(filename string) (wf *File, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
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
func ReadFloat64(filename string) (data []float64, rate uint32, err error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return
	}
	rate = wf.SampleRate
	data, err = wf.ToFloat64(0, 0)
	return
}

// ReadFloat32 reads a wav file and returns the samples as float32
func ReadFloat32(filename string) (data []float32, rate uint32, err error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return
	}
	rate = wf.SampleRate
	data, err = wf.ToFloat32(0, 0)
	return
}

// ReadInt16 reads a wav file and returns the samples as int16
func ReadInt16(filename string) (data []int16, rate uint32, err error) {
	wf, err := ReadFile(filename)
	if err != nil {
		return
	}
	rate = wf.SampleRate
	data, err = wf.ToInt16(0, 0)
	return
}
