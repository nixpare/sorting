package sorting

import (
	"testing"
)

func TestTimSort(t *testing.T) {
	testSortingAlgorithmStandard(t, TimSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkTimSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, TimSort, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, TimSort, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, TimSort, inOrderInteger, testSize)
	})
}
