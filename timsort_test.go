package sorting

import (
	"testing"
)

func TestTimSort(t *testing.T) {
	testSortingAlgorithmStandard(t, TimSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkTimSort(b *testing.B) {
	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard(b, TimSort, newRandomInteger)
	})

	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced(b, TimSort, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle(b, TimSort, inOrderInteger, testSize)
	})
}
