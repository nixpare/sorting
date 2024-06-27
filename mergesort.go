package sorting

import (
	"sync"
)

var (
	MergeSortMaxSortingThreads = 64
	MergeSortMinSliceLength    = 32
)

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

	mergeExternalBuffer(v, mid, tmp)
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

	mergeExternalBuffer(v, mid, tmp)
}

func mergeExternalBuffer[T Comparable[T]](v []T, mid int, tmp []T) {
	if v[mid-1].Compare(v[mid]) < 0 {
		return
	}

	if v[0].Compare(v[len(v)-1]) > 0 {
		swap(v, mid, tmp)
		return
	}
	
	if mid <= len(v) - mid {
		mergeExternalBufferFromLeft(v, mid, tmp)
	} else {
		mergeExternalBufferFromRight(v, mid, tmp)
	}
}

func mergeExternalBufferFromLeft[T Comparable[T]](v []T, mid int, tmp []T) {
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
		i++
		k++
	}

	for j < len(v) {
		v[k] = v[j]
		j++
		k++
	}
}

func mergeExternalBufferFromRight[T Comparable[T]](v []T, mid int, tmp []T) {
	n := copy(tmp, v[mid:])
	tmp = tmp[:n]

	i, j, k := mid-1, len(tmp)-1, len(v)-1
	for i >= 0 && j >= 0 {
		if v[i].Compare(tmp[j]) > 0 {
			v[k] = v[i]
			i--
		} else {
			v[k] = tmp[j]
			j--
		}

		k--
	}

	for i >= 0 {
		v[k] = v[i]
		i--
		k--
	}

	for j >= 0 {
		v[k] = tmp[j]
		j--
		k--
	}
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

	mergeInternalBuffer(v, buffer, mid)
	mergeSortUnstable(v[:buffer], tmp)

	mergeExternalBuffer(v, buffer, tmp)
}

func mergeInternalBuffer[T Comparable[T]](v []T, buffer int, right int) {
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
		left++
		buffer++
	}
}

func swap[T Comparable[T]](v []T, mid int, tmp []T) {
	if mid <= len(v) - mid {
		n := copy(tmp, v[:mid])
		
		i := 0
		for j := mid; j < len(v); j++ {
			v[i], v[j] = v[j], v[i]
			i++
		}

		copy(v[i:], tmp[:n])
	} else {
		n := copy(tmp, v[mid:])
		
		i := len(v)-1
		for j := mid-1; j >= 0; j-- {
			v[i], v[j] = v[j], v[i]
			i--
		}

		copy(v[:i], tmp[:n])
	}
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
