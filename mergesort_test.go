package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	TestSortingAlgorithm(t, MergeSort, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, MergeSort, newRandomInteger)
}

func TestMergeSortMulti(t *testing.T) {
	TestSortingAlgorithm(t, MergeSortMulti, newRandomInteger, integerSliceIsSorted)
}

func BenchmarkMergeSortMulti(b *testing.B) {
	BenchmarkSortingAlgorithm(b, MergeSortMulti, newRandomInteger)
}
