package game

import (
    "testing"
)
var binaryCompareInt64Datasets = []struct {
	in  [2]int64
	out bool
}{
	{[2]int64{34905131008, 545392672}, false},
	{[2]int64{34905131008, 34905131008}, true},
}

func TestBinaryCompareInt64(t *testing.T) {

	for _, data := range binaryCompareInt64Datasets {
		result := BinaryCompareInt64(data.in[0], data.in[1])

		if result != data.out {
			t.Errorf("Error : got %t, want %t", result, data.out)
		}
	}
}
