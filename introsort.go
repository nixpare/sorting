package sorting

import "math"

// IntroSort is a not-stable algorithm which mixes InsertionSort, HeapSort
// and DoublePivotQuickSort. Not recommended for slices nearly sorted.
func IntroSort[S ~[]E, E Ordered[E]](v S) {
	maxDepth := 2 * int(math.Round(math.Log2(float64(len(v)))))
	introSort(v, maxDepth)
}

func introSort[S ~[]E, E Ordered[E]](v S, maxDepth int) {
	if len(v) < InsertionSortSliceLength {
		InsertionSort(v)
		return
	}

	if maxDepth == 0 {
		HeapSort(v)
		return
	}

	p, q := partition(v)

    introSort(v[:p], maxDepth-1)
    introSort(v[p+1:q], maxDepth-1)
    introSort(v[q+1:], maxDepth-1)
}