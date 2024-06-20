package sorting

import (
	"testing"
)

func TestMergeSortStandard(t *testing.T) {
	testSortingAlgorithm(t, MergeSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSortStandard(b *testing.B) {
	benchmarkSortingAlgorithm(b, MergeSort, newRandomInteger)
}

func TestMergeSortMulti(t *testing.T) {
	testSortingAlgorithm(t, MergeSortMulti, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSortMulti(b *testing.B) {
	benchmarkSortingAlgorithm(b, MergeSortMulti, newRandomInteger)
}

func TestMergeSortUnstable(t *testing.T) {
	testSortingAlgorithm(t, MergeSortUnstable, newRandomInteger, nil)
}

func BenchmarkMergeSortUnstable(b *testing.B) {
	benchmarkSortingAlgorithm(b, MergeSortUnstable, newRandomInteger)
}
