package sorting

import (
	"testing"
)

func TestTimSort(t *testing.T) {
	testSortingAlgorithm(t, TimSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkTimSort(b *testing.B) {
	benchmarkSortingAlgorithm(b, TimSort, newRandomInteger)
}
