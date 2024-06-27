package sorting

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortingAlgorithmStandard[[]integer](t, HeapSort, newRandomInteger, nil)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSortingAlgorithmStandard[[]integer](b, HeapSort, newRandomInteger)
}
