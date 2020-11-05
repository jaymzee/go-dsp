package wav

// ReadFloat64 reads the wav file data and returns it as float64
func ReadFloat64(filename string) ([]float64, error) {
	w, err := ReadWave(filename)
	if err != nil {
		return nil, err
	}
	return w.DataFloat64()
}

// ReadFloat32 reads the wav file data and returns it as float32
func ReadFloat32(filename string) ([]float32, error) {
	w, err := ReadWave(filename)
	if err != nil {
		return nil, err
	}
	return w.DataFloat32()
}
