package sorting

import (
	"slices"
)

var (
	TimSortRun = 32
)

type sliceRange struct {
	from int
	to   int
}

// TimSort is a simple implementation of the original algorithm, it is stable and not in-place.
// It's a work in progress: right now, considering it uses the same MergeExternal function used in
// MergeSort, it has practically the same performance of said algorithm, with the only difference that
// it does not uses recursion. However MergeSort still performs slightly better.
func TimSort[T Comparable[T]](v []T) {
	if len(v) < 2 {
		return
	}

	var runs []sliceRange

	prev, ascending := 0, v[0].Compare(v[1]) <= 0
	// Sort individual subarrays of size RUN
	for i := 1; i < len(v); {
		if v[i].Compare(v[i-1]) <= 0 {
			if ascending {
				i++
				continue
			}

			if i-prev >= TimSortRun {
				slices.Reverse(v[prev:i])
				runs = append(runs, sliceRange{from: prev, to: i})
				prev = i
			} else {
				to := min(prev+TimSortRun, len(v))

				InsertionSort(v[prev:to])
				runs = append(runs, sliceRange{from: prev, to: to})

				prev = to
			}
		} else {
			if !ascending {
				i++
				continue
			}

			if i-prev >= TimSortRun {
				runs = append(runs, sliceRange{from: prev, to: i})
				prev = i
			} else {
				to := min(prev+TimSortRun, len(v))

				InsertionSort(v[prev:to])
				runs = append(runs, sliceRange{from: prev, to: to})

				prev = to
			}
		}

		if prev == len(v) {
			break
		}

		if prev == len(v)-1 {
			runs = append(runs, sliceRange{from: prev, to: prev + 1})
			break
		}

		ascending = v[prev].Compare(v[prev+1]) <= 0
		i = prev + 1
	}

	tmp := make([]T, len(v))

	for j := 0; len(runs) > 1; j++ {
		i := 0
		for ; i < len(runs)/2; i++ {
			run1 := runs[i*2]
			run2 := runs[i*2+1]
			if run1.to != run2.from {
				panic("not contiguos runs detected")
			}

			MergeExternal(v[run1.from:run2.to], run1.to-run1.from, tmp)
			runs[i] = sliceRange{from: run1.from, to: run2.to}
		}

		if i*2 == len(runs)-1 {
			lastRun := runs[len(runs)-1]
			runs[i] = sliceRange{from: lastRun.from, to: lastRun.to}
			i++
		}

		runs = runs[:i]
	}
}
