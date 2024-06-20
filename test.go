package sorting

import (
	"math/rand"
	"testing"
)

const (
	testTimes = 20
	testSize  = 1_000_000
)

type data struct {
	x   int
	pos int
}

func(d data) Compare(other data) int {
	if d.x == other.x {
		return d.pos - other.pos
	}

	return d.x - other.x
}

type integer data

func (i integer) Compare(other integer) int {
	return i.x - other.x
}

func newRandomInteger(i int) integer {
	x := rand.Intn(testSize)
	return integer(data{ x: x, pos: i })
}

func integerSliceIsSorted(v []integer) bool {
	for i := 1; i < len(v); i++ {
		if data(v[i]).Compare(data(v[i-1])) < 0 {
			return false
		}
	}

	return true
}

func TestSortingAlgorithm[T Comparable[T]](t *testing.T, algo func(v []T), randGenFunc func(i int) T, isSortedFunc func([]T) bool) {
	t.Helper()
	v := make([]T, testSize)

	for range testTimes {
		for i := range v {
			v[i] = randGenFunc(i)
		}

		algo(v)

		if !IsSorted(v) {
			t.Error("not sorted", v)
		}
	}
}

func BenchmarkSortingAlgorithm[T Comparable[T]](b *testing.B, algo func(v []T), randGenFunc func(i int) T) {
	b.Helper()
	b.ReportAllocs()

	v := make([]T, testSize)

	for range b.N {
		b.StopTimer()
		for i := range v {
			v[i] = randGenFunc(i)
		}
		b.StartTimer()

		algo(v)
	}
}
