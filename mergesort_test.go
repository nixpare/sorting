package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testSortingAlgorithmStandard(t, MergeSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, MergeSort, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, MergeSort, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, MergeSort, inOrderInteger, testSize)
	})
}

func TestMergeSortMulti(t *testing.T) {
	testSortingAlgorithmStandard(t, MergeSortMulti, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSortMulti(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, MergeSortMulti, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, MergeSortMulti, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, MergeSortMulti, inOrderInteger, testSize)
	})
}

func TestMergeSortUnstable(t *testing.T) {
	testSortingAlgorithmStandard[[]integer](t, MergeSortUnstable, newRandomInteger, nil)
}

func BenchmarkMergeSortUnstable(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, MergeSortUnstable, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, MergeSortUnstable, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, MergeSortUnstable, inOrderInteger, testSize)
	})
}
