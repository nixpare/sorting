package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testSortingAlgorithmStandard(t, MergeSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard(b, MergeSort, newRandomInteger)
	})

	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced(b, MergeSort, newRandomInteger)
	})
}

func TestMergeSortMulti(t *testing.T) {
	testSortingAlgorithmStandard(t, MergeSortMulti, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSortMulti(b *testing.B) {
	benchmarkSortingAlgorithmStandard(b, MergeSortMulti, newRandomInteger)
}

func TestMergeSortUnstable(t *testing.T) {
	testSortingAlgorithmStandard(t, MergeSortUnstable, newRandomInteger, nil)
}

func BenchmarkMergeSortUnstable(b *testing.B) {
	benchmarkSortingAlgorithmStandard(b, MergeSortUnstable, newRandomInteger)
}
