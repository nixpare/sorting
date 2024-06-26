package sorting

import (
	"testing"
)

func TestTimSort(t *testing.T) {
	testSortingAlgorithmStandard(t, TimSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkTimSort(b *testing.B) {
	benchmarkSortingAlgorithmStandard(b, TimSort, newRandomInteger)
}
