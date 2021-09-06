package signal

import (
	"testing"
)


func TestMax1(t *testing.T) {
	got := Max(-3.1, 2.2, 3.3)
	if got != 3.3 {
		t.Errorf("Max(-3.1, 2.2, 3.3) = %v; want 3.3", got)
	}
}

func TestMax2(t *testing.T) {
	got := Max([]float64{-3.1, 2.2, 3.3}...)
	if got != 3.3 {
		t.Errorf("Max(-3.1, 2.2, 3.3) = %v; want 3.3", got)
	}
}

func TestMin1(t *testing.T) {
	got := Min(-3.1, 2.2, 3.3)
	if got != -3.1 {
		t.Errorf("Min(-3.1, 2.2, 3.3) = %v; want -3.1", got)
	}
}

func TestLog2(t *testing.T) {
	table := map[int]int{
		1:  0,
		2:  1,
		4:  2,
		8:  3,
		16: 4,
	}
	for x, y := range table {
		got := Log2(x)
		if got != y {
			t.Errorf("log2(%d) = %d; want %d", x, got, y)
		}
	}
}
