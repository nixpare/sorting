package sorting

// InsertionSort is a standard implementation of the algorithm
func InsertionSort[S ~[]E, E Ordered[E]](v S) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0 && v[j].Compare(v[j-1]) < 0; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}
