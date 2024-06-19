package sorting

import (
	"testing"
)

const (
	testTimes = 1
	testSize  = 1_000_000
)

type integer int

func (i integer) Compare(other integer) int {
	return int(i) - int(other)
}

func TestSortingAlgorithm[T Comparable[T]](t *testing.T, algo func(v []T), randGen func() T) {
	t.Helper()
	v := make([]T, testSize)

	for range testTimes {
		for i := range v {
			v[i] = randGen()
		}

		algo(v)

		if !IsSorted(v) {
			t.Error("not sorted", v)
		}
	}
}

func BenchmarkSortingAlgorithm[T Comparable[T]](b *testing.B, algo func(v []T), randGen func() T) {
	b.Helper()
	b.ReportAllocs()

	v := make([]T, testSize)

	for range b.N {
		b.StopTimer()
		for i := range v {
			v[i] = randGen()
		}
		b.StartTimer()

		algo(v)
	}
}
