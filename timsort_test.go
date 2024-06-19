package sorting

import (
	"math/rand"
	"testing"
)

func TestTimSort(t *testing.T) {
	TestSortingAlgorithm(t, TimSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}

func BenchmarkTimSort(b *testing.B) {
	BenchmarkSortingAlgorithm(b, TimSort, func() integer {
		return integer(rand.Intn(testSize))
	})
}
