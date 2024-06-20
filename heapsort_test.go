package sorting

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortingAlgorithm(t, HeapSort, newRandomInteger, nil)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSortingAlgorithm(b, HeapSort, newRandomInteger)
}
