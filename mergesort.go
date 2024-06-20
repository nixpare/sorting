package sorting

import (
	"sync"
)

func MergeSort[T Comparable[T]](v []T) {
	mergeSort(v, newBuffer(v))
}

func mergeSort[T Comparable[T]](v []T, tmp []T) {
	switch len(v) {
	case 0, 1:
		return
	case 2:
		if v[0].Compare(v[1]) > 0 {
			v[0], v[1] = v[1], v[0]
		}
		return
	case 3:
		insertionSort(v)
		return
	}
	
	mid := len(v) / 2
	if mid % 2 == 1 {
		mid++
	}
	buffer := mid / 2

	mergeSort(v[mid:], tmp)        // Sort right part
	mergeSort(v[:buffer], tmp)     // Sort half of the left part

	left := mergeInternalBuffer(v, buffer, mid)  // Merges the left half part with the right full part to the right
	mergeSort(v[:left], tmp)       // Sort recursively the remaining left half part

	mergeExternalBuffer(v, left, tmp)      // Merges the 1/4 to the left with the 3/4 to the right
}

func mergeInternalBuffer[T Comparable[T]](v []T, buffer int, right int) int {
	left := 0

	for buffer < right && right < len(v) {
		if v[left].Compare(v[right]) <= 0 {
			v[buffer], v[left] = v[left], v[buffer]
			left++
		} else {
			v[buffer], v[right] = v[right], v[buffer]
			right++
		}

		buffer++
	}

	for buffer < right {
		v[buffer], v[left] = v[left], v[buffer]
		left++; buffer++
	}

	return left
}

func mergeExternalBuffer[T Comparable[T]](v []T, mid int, tmp []T) {
	copy(tmp, v[:mid])

	i, j, k := 0, mid, 0
	for i < mid && j < len(v) {
		if tmp[i].Compare(v[j]) <= 0 {
			v[k] = tmp[i]
			i++
		} else {
			v[k] = v[j]
			j++
		}

		k++
	}

	for i < mid {
		v[k] = tmp[i]
		i++; k++
	}

	for j < len(v) {
		v[k] = v[j]
		j++; k++
	}
}

/* func mergeInPlace[T Comparable[T]](v []T, mid int) {
	
} */

func newBuffer[T any](v []T) []T {
	n := len(v) / 2
	if n % 2 == 1 {
		n++
	}

	n /= 2
	return make([]T, n) 
}

var (
	MaxSortingThreads = 64
)

func MergeSortMulti[T Comparable[T]](v []T) {
	mergeSortMulti(v, make([]T, len(v)), 1)
}

func mergeSortMulti[T Comparable[T]](v []T, tmp []T, threads int) {
	switch len(v) {
	case 0, 1:
		return
	case 2:
		if v[0].Compare(v[1]) > 0 {
			v[0], v[1] = v[1], v[0]
		}
		return
	case 3:
		insertionSort(v)
		return
	}
	
	mid := len(v) / 2
	if mid % 2 == 1 {
		mid++
	}
	buffer := mid / 2

	if threads < MaxSortingThreads {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			mergeSortMulti(v[mid:], tmp[mid:], threads*2)          // Sort right part
			wg.Done()
		}()

		go func() {
			mergeSortMulti(v[:buffer], tmp[:buffer], threads*2)   // Sort half of the left part
			wg.Done()
		}()

		wg.Wait()
	} else {
		mergeSort(v[mid:], tmp[mid:])         // Sort right part
		mergeSort(v[:buffer], tmp[:buffer])   // Sort half of the left part
	}

	left := mergeInternalBuffer(v, buffer, mid)                   // Merges the left half part with the right full part to the right
	mergeSortMulti(v[:left], tmp[:left], threads)   // Sort recursively the remaining left half part

	mergeExternalBuffer(v, left, tmp)   // Merges the 1/4 to the left with the 3/4 to the right
}
