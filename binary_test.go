package sorting

import "testing"

func TestBinarySearch(t *testing.T) {
	v := integerSlice([]int{ 1, 2, 3, 4, 5, 6, 8, 9, 10 })

	type testData struct{ input integer; output int }
	tests := []testData{
		{ input: integer(data{ x: 0 }), output: -1 },
		{ input: integer(data{ x: 7 }), output: -1 },
		{ input: integer(data{ x: 11 }), output: -1 },
	}

	for i, x := range v {
		if idx := BinarySearch(v, x); idx != i {
			t.Errorf("binary search failed: expected %d but found %d", i, idx)
		}
	}
	for _, test := range tests {
		if idx := BinarySearch(v, test.input); idx != test.output {
			t.Errorf("binary search failed: expected %d but found %d", test.output, idx)
		}
	}
}
