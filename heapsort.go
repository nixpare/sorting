package sorting

func HeapSort[T Comparable[T]](v []T) {
	createHeap(v)
	for i := len(v)-1; i > 0; i-- {
		v[0], v[i] = v[i], v[0]
		heapify(v, 0, i)
	}
}

func createHeap[T Comparable[T]](v []T) {
	for i := len(v) / 2; i >= 0; i-- {
		heapify(v, i, len(v))
	}
}

func heapify[T Comparable[T]](v []T, i int, n int) {
	x := v[i]

	for {
		j := i * 2 + 1
		if j >= n {
			break
		}

		if j+1 < n && v[j+1].Compare(v[j]) > 0 {
			j++
		}

		if v[j].Compare(x) <= 0 {
			break
		}

		v[i] = v[j]
		i = j
	}

	v[i] = x
}