package sorting

func insertionSort[T Comparable[T]](v []T) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0 && v[j].Compare(v[j-1]) < 0; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}
