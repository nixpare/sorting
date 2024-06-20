package sorting

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	TestSortingAlgorithm(t, HeapSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkHeapSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, HeapSort, newRandomInteger)
}
