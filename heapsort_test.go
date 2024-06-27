package sorting

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortingAlgorithmStandard[[]integer](t, HeapSort, newRandomInteger, nil)
}

func BenchmarkHeapSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, HeapSort, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, HeapSort, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, HeapSort, inOrderInteger, testSize)
	})
}
