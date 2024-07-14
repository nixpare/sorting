package sorting

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

const (
	testTimes        = 20
	testSize         = 1_000_000
	reducedTestTimes = 100
	reducedTestSize  = 64
)

type data struct {
	x   int
	pos int
}

func (d data) Compare(other data) int {
	if d.x == other.x {
		return d.pos - other.pos
	}

	return d.x - other.x
}

func (d data) String() string {
	return fmt.Sprintf("%d:%d", d.x, d.pos)
}

type integer data

func (i integer) Compare(other integer) int {
	return i.x - other.x
}

func (i integer) String() string {
	return data(i).String()
}

func newRandomInteger(i int) integer {
	x := rand.Intn(testSize)
	return integer(data{x: x, pos: i})
}

func inOrderInteger(i int) integer {
	return integer(data{x: i, pos: i})
}

func integerSlice(v []int) []integer {
	u := make([]integer, len(v))
	for i := range v {
		u[i] = integer(data{x: v[i], pos: i})
	}
	return u
}

func integerSliceIsSorted(v []integer) bool {
	for i := 1; i < len(v); i++ {
		if data(v[i]).Compare(data(v[i-1])) < 0 {
			return false
		}
	}

	return true
}

func testSortingAlgorithmStandard[S ~[]E, E Ordered[E]](t *testing.T, algo func(v S), randGenFunc func(i int) E, isStableSortedFunc func(S) bool) {
	t.Helper()
	testSortingAlgorithm(t, algo, randGenFunc, isStableSortedFunc, testTimes, testSize)
}

func testSortingAlgorithmReduced[S ~[]E, E Ordered[E]](t *testing.T, algo func(v S), randGenFunc func(i int) E, isStableSortedFunc func(S) bool) {
	t.Helper()
	testSortingAlgorithm(t, algo, randGenFunc, isStableSortedFunc, reducedTestTimes, reducedTestSize)
}

func testSortingAlgorithm[S ~[]E, E Ordered[E]](
	t *testing.T, algo func(v S),
	randGenFunc func(i int) E, isStableSortedFunc func(S) bool,
	times int, size int,
) {
	v := make([]E, size)

	for range times {
		for i := range v {
			v[i] = randGenFunc(i)
		}

		algo(v)

		if isStableSortedFunc != nil {
			if !isStableSortedFunc(v) {
				t.Error("not sorted")
			}
		} else {
			if !IsSorted(v) {
				t.Error("not sorted")
			}
		}
	}
}

func benchmarkSortingAlgorithmStandard[S ~[]E, E Ordered[E]](b *testing.B, algo func(v S), randGenFunc func(i int) E) {
	b.Helper()
	benchmarkSortingAlgorithm(b, algo, randGenFunc, testSize)
}

func benchmarkSortingAlgorithmReduced[S ~[]E, E Ordered[E]](b *testing.B, algo func(v S), randGenFunc func(i int) E) {
	b.Helper()
	benchmarkSortingAlgorithm(b, algo, randGenFunc, reducedTestSize)
}

func benchmarkSortingAlgorithm[S ~[]E, E Ordered[E]](b *testing.B, algo func(v S), randGenFunc func(i int) E, size int) {
	v := make([]E, size)

	for range b.N {
		b.StopTimer()
		for i := range v {
			v[i] = randGenFunc(i)
		}
		b.StartTimer()

		algo(v)
	}
}

var shuffles = [...]int{ 90, 60, 20, 5, 0 }

func benchmarkSortingAlgorithmShuffle[S ~[]E, E Ordered[E]](b *testing.B, algo func(v S), inOrderFunc func(i int) E, size int) {
	v := make([]E, size)
	for i := range v {
		v[i] = inOrderFunc(i)
	}

	for _, shuffle := range shuffles {
		times := size / 100 * shuffle

		b.Run(fmt.Sprintf("%d%%", shuffle), func(b *testing.B) {
			for range b.N {
				b.StopTimer()

				prev := rand.Intn(size)
				for range times {
					next := rand.Intn(size)
					v[prev], v[next] = v[next], v[prev]
					prev = next
				}
				b.StartTimer()

				algo(v)
			}
		})
	}
}

func BenchmarkGoDefaultStable(b *testing.B) {
	algo := func(v []integer) {
		slices.SortStableFunc(v, func(a, b integer) int {
			return a.Compare(b)
		})
	}

	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, algo, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, algo, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, algo, inOrderInteger, testSize)
	})
}

func BenchmarkGoDefault(b *testing.B) {
	algo := func(v []integer) {
		slices.SortFunc(v, func(a, b integer) int {
			return a.Compare(b)
		})
	}

	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, algo, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, algo, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, algo, inOrderInteger, testSize)
	})
}
