package signal

import "testing"

func TestConv1(t *testing.T) {
	got := Conv([]float64{1, 2, 3}, []float64{4, 5})
	if !Equal(got, []float64{4, 13, 22, 15}) {
		t.Errorf("conv([1 2 3], [4 5]) = %v; want [4 13 22 15]", got)
	}
}

func TestConv2(t *testing.T) {
	got := Conv([]float64{1, 2, 3}, []float64{4, 5, 6})
	if !Equal(got, []float64{4, 13, 28, 27, 18}) {
		t.Errorf("conv([1 2 3], [4 5 6]) = %v; want [4 13 28 27 18]", got)
	}
}

func TestConv3(t *testing.T) {
	got := Conv([]float64{2}, []float64{3, 4, 5})
	if !Equal(got, []float64{6, 8, 10}) {
		t.Errorf("conv([2], [3 4 5]) = %v; want [6 8 10]", got)
	}
}

func TestConv4(t *testing.T) {
	got := Conv([]float64{2, 3, 4, 5}, []float64{6, 7})
	if !Equal(got, []float64{12, 32, 45, 58, 35}) {
		t.Errorf("conv([2 3 4 5], [6 7]) = %v; want [12 32 45 58 35]", got)
	}
}

func Equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var (
        x6 = []float64{1, 2, 3, 4, 5, 6}
        h6 = []float64{4, 5, 6, 7, 8, 9}
        x12 = []float64{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 2, 1}
        h12 = []float64{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 2, 1}
        x24 [24]float64
        h24 [24]float64
        x48 [48]float64
        h48 [48]float64
)

func BenchmarkConv6(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Conv(x6, h6)
        }
}

func BenchmarkConv12(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Conv(x12, h12)
        }
}

func BenchmarkConv24(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Conv(x24[:], h24[:])
        }
}

func BenchmarkConv48(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Conv(x48[:], h48[:])
        }
}
