package sorting

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	TestSortingAlgorithm(t, MergeSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}

func BenchmarkMergeSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, MergeSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}
