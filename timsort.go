package sorting

import (
	"slices"
)

var (
	TimSortRun = 32
)

// TimSort is a simple implementation of the original algorithm, it is stable and not in-place.
// It's a work in progress: right now, considering it uses the same MergeExternal function used in
// MergeSort, it has practically the same performance of said algorithm, with the only difference that
// it does not uses recursion. However MergeSort still performs slightly better.
func TimSort[S ~[]E, E Ordered[E]](v S) {
	if len(v) < 2 {
		return
	}

	var runs []int

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
				runs = append(runs, i)
				prev = i
			} else {
				to := min(prev+TimSortRun, len(v))

				InsertionSort(v[prev:to])
				runs = append(runs, to)
				prev = to
			}
		} else {
			if !ascending {
				i++
				continue
			}

			if i-prev >= TimSortRun {
				runs = append(runs, i)
				prev = i
			} else {
				to := min(prev+TimSortRun, len(v))

				InsertionSort(v[prev:to])
				runs = append(runs, to)
				prev = to
			}
		}

		if prev == len(v) {
			break
		}

		if prev == len(v)-1 {
			runs = append(runs, len(v))
			break
		}

		ascending = v[prev].Compare(v[prev+1]) <= 0
		i = prev + 1
	}

	tmp := make([]E, len(v))

	for j := 0; len(runs) > 1; j++ {
		i, from := 0, 0
		for ; i < len(runs)/2; i++ {
			mid := runs[i*2]
			to := runs[i*2+1]

			MergeExternal(v[from:to], mid - from, tmp)
			runs[i], from = to, to
		}

		if i*2 == len(runs)-1 {
			runs[i] = len(v)
			i++
		}

		runs = runs[:i]
	}
}
