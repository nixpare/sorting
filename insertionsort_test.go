package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	t.Run("Reduced", func(t *testing.T) {
		testSortingAlgorithmReduced(t, InsertionSort, newRandomInteger, integerSliceIsSorted)
	})
}

func BenchmarkInsertionSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, InsertionSort, newRandomInteger)
	})
}
