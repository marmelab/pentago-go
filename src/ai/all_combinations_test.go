package ai

import (
    "testing"
)
var countBitsForCombinationIfStillPossibleDataset = []struct {
	in  [3]int64
	out int
}{
	{[3]int64{0b10000_100000_100000_100000_100000_000000, 0b10000_100000_100000_000000_100000_100011, 0b00000_000000_000000_100000_000000_000000}, 0},
	{[3]int64{0b10000_100000_100000_100000_100000_000000, 0b10000_100000_100000_000000_100000_100011, 0b00000_000000_000000_010000_000000_100001}, 4},
}
func TestCountBitsForCombinationIfStillPossible(t *testing.T) {
	for _, data := range countBitsForCombinationIfStillPossibleDataset {
		result := CountBitsForCombinationIfStillPossible(data.in[0], data.in[1], data.in[2])
		expected_result := data.out
	
		if result != expected_result {
			t.Errorf("Error CountBitsForCombinationIfStillPossible : got %d, want %d", result, expected_result)
		}
	}


}

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

var countCommonBitsBetweenTwoInt64Dataset = []struct {
	in [2]int64
	out int
}{
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b000000_000000_001111_000000_000000_000000}, 0},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_000000_001111_000000_000000_000000}, 1},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_100000_001111_000000_000000_000000}, 2},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_100000_101111_000000_000000_000000}, 3},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_100000_101111_000000_100000_000000}, 4},
	{[2]int64{0b100000_100000_100000_100000_100000_000000, 0b100000_100000_101111_100000_100000_000000}, 5},
}
func TestCountCommonBitsBetweenTwoInt64(t *testing.T) {
	for _, data := range countCommonBitsBetweenTwoInt64Dataset {
		result := CountCommonBitsBetweenTwoInt64(data.in[0], data.in[1])

		if result != data.out {
			t.Errorf("Error CountCommonBitsBetweenTwoInt64 : got %d, want %d", result, data.out)
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
