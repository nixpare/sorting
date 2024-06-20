package sorting

import (
	"testing"
)

func TestTimSort(t *testing.T) {
	TestSortingAlgorithm(t, TimSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkTimSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, TimSort, newRandomInteger)
}
