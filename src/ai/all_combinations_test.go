package ai

import (
    "testing"
)
var binaryCompareInt64Datasets = []struct {
	in  [2]int64
	out bool
}{
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_100000_001111_100000_100000_000000}, false},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b110000_110000_110000_101000_110100_010000}, true},
}

func TestBinaryCompareInt64(t *testing.T) {

	for _, data := range binaryCompareInt64Datasets {
		result := BinaryCompareInt64(data.in[0], data.in[1])

		if result != data.out {
			t.Errorf("Error : got %t, want %t", result, data.out)
		}
	}
}

var isBoardFullDataset = []struct {
	in  [2]int64
	out bool
}{
	{[2]int64{0b111111_111111_111111_111111_111111_111111, 0b110111_111111_111111_111111_111111_111111}, false},
	{[2]int64{0b111111_111111_111111_111111_111111_111111, 0b111111_111111_111111_111111_111111_111111}, true},
}

func TestIsBoardFull(t *testing.T) {
	for _, data := range isBoardFullDataset {
		result := BinaryCompareInt64(data.in[0], data.in[1])

		if result != data.out {
			t.Errorf("Error : got %t, want %t", result, data.out)
		}
	}
}
