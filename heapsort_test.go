package sorting

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortingAlgorithmStandard(t, HeapSort, newRandomInteger, nil)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSortingAlgorithmStandard(b, HeapSort, newRandomInteger)
}
