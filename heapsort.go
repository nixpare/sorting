package sorting

// HeapSort is a standard implementation of the algorithm, without using
// additional space.
func HeapSort[S ~[]E, E Ordered[E]](v S) {
	createHeap(v)
	for i := len(v) - 1; i > 0; i-- {
		v[0], v[i] = v[i], v[0]
		heapify(v, 0, i)
	}
}

func createHeap[S ~[]E, E Ordered[E]](v S) {
	for i := len(v) / 2; i >= 0; i-- {
		heapify(v, i, len(v))
	}
}

func heapify[S ~[]E, E Ordered[E]](v S, i int, n int) {
	x := v[i]

	for {
		j := i*2 + 1
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
