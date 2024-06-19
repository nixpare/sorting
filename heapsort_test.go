package sorting

import (
	"math/rand"
	"testing"
)

func TestHeapSort(t *testing.T) {
	TestSortingAlgorithm(t, HeapSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}

func BenchmarkHeapSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, HeapSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}
