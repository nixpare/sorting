package sorting

// DoublePivotQuickSort implements the quicksort algorithm in its optimized
// version, but it's still not recommendable compared to MergeSort or TimSort
// if the slice is nearly sorted.
func DoublePivotQuickSort[S ~[]E, E Ordered[E]](v S) {
	if len(v) < 2 {
		return
	}

	p, q := partition(v)

    DoublePivotQuickSort(v[:p])
    DoublePivotQuickSort(v[p+1:q])
    DoublePivotQuickSort(v[q+1:])
}

func partition[S ~[]E, E Ordered[E]](v S) (int, int) {
	if v[0].Compare(v[len(v)-1]) > 0 {
		v[0], v[len(v)-1] = v[len(v)-1], v[0]
	}
	p, q := v[0], v[len(v)-1]

	i, j, k := 1, 1, len(v)-2

	for j <= k {
		if v[j].Compare(p) < 0 {
			v[i], v[j] = v[j], v[i]
			i++
			j++
		} else if v[j].Compare(q) < 0 {
			j++
		} else {
			v[j], v[k] = v[k], v[j]
			k--
		}
	}
	i--
	k++
	v[0], v[i] = v[i], v[0]
	v[len(v)-1], v[k] = v[k], v[len(v)-1]
	return i, k
}
