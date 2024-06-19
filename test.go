package sorting

import (
	"slices"
	"testing"
)

const (
	testTimes = 20
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

		if !slices.IsSortedFunc(v, func(a, b T) int {
			return a.Compare(b)
		}) {
			t.Error("not sorted")
		}
	}
}

func BenchmarkSortingAlgorithm[T Comparable[T]](b *testing.B, algo func(v []T), randGen func() T) {
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