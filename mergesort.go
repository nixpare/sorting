package sorting

import (
	"sync"
)

var (
	MergeSortMaxSortingThreads = 64
	MergeSortMinSliceLength    = 32
)

// MergeSort is based on the traditional algorithm with a few optimizations:
//   - the external buffer for the merge is half the slice length
//   - for sub-slices shorter than MergeSortMinSliceLength it uses InsertionSort
//   - other optimizations to the MergeExternal algorithm
// For a multithreaded version see MergeSortMulti
func MergeSort[T Comparable[T]](v []T) {
	if len(v) <= MergeSortMinSliceLength {
		mergeSort(v, nil)
	} else {
		mergeSort(v, newBuffer(v, 1))
	}
}

func MergeSortMulti[T Comparable[T]](v []T) {
	mergeSortMulti(v, newBuffer(v, 0), 1)
}

func MergeSortUnstable[T Comparable[T]](v []T) {
	mergeSortUnstable(v, newBuffer(v, 2))
}

func mergeSort[T Comparable[T]](v []T, tmp []T) {
	if len(v) < 2 {
		return
	}

	if len(v) <= MergeSortMinSliceLength {
		InsertionSort(v)
		return
	}

	mid := len(v) / 2

	mergeSort(v[:mid], tmp)
	mergeSort(v[mid:], tmp)

	MergeExternal(v, mid, tmp)
}

func mergeSortMulti[T Comparable[T]](v []T, tmp []T, threads int) {
	if len(v) < 2 {
		return
	}

	if len(v) <= MergeSortMinSliceLength {
		InsertionSort(v)
		return
	}

	mid := len(v) / 2

	if threads < MergeSortMaxSortingThreads {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			mergeSortMulti(v[:mid], tmp[:mid], threads*2)
			wg.Done()
		}()

		go func() {
			mergeSortMulti(v[mid:], tmp[mid:], threads*2)
			wg.Done()
		}()

		wg.Wait()
	} else {
		mergeSort(v[:mid], tmp)
		mergeSort(v[mid:], tmp)
	}

	MergeExternal(v, mid, tmp)
}

func mergeSortUnstable[T Comparable[T]](v []T, tmp []T) {
	if len(v) < 2 {
		return
	}

	if len(v) <= MergeSortMinSliceLength {
		InsertionSort(v)
		return
	}

	mid := len(v) / 2
	if mid%2 == 1 {
		mid++
	}
	buffer := mid / 2

	mergeSortUnstable(v[mid:], tmp)
	mergeSortUnstable(v[:buffer], tmp)

	MergeInternal(v, buffer, mid)
	mergeSortUnstable(v[:buffer], tmp)

	MergeExternal(v, buffer, tmp)
}

func newBuffer[T any](v []T, splitTimes int) []T {
	if splitTimes == 0 {
		return make([]T, len(v))
	}

	n := len(v) / 2

	for range splitTimes - 1 {
		if n%2 == 1 {
			n++
		}

		n /= 2
	}

	return make([]T, n)
}
