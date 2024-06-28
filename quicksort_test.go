package sorting

import "testing"

func TestDoublePivotQuickSort(t *testing.T) {
	testSortingAlgorithmStandard[[]integer](t, DoublePivotQuickSort, newRandomInteger, nil)
}

func BenchmarkDoublePivotQuickSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, DoublePivotQuickSort, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, DoublePivotQuickSort, newRandomInteger)
	})

	b.Run("Shuffle_200.000", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, DoublePivotQuickSort, inOrderInteger, 200_000)
	})
}