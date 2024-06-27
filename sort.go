// A collection of modified sorting algorithms for Go
package sorting

type Ordered[E any] interface {
	Compare(E) int
}

// Sort is the best stable sorting algorithm of the package.
// In this version, this algorithm is [sorting.MergeSort]
func Sort[S ~[]E, E Ordered[E]](v S) {
	MergeSort(v)
}

// SortMulti is the best stable multithreaded sorting algorithm of the package.
// In this version, this algorithm is [sorting.MergeSortMulti]
func SortMulti[S ~[]E, E Ordered[E]](v S) {
	MergeSortMulti(v)
}

func IsSorted[S ~[]E, E Ordered[E]](v S) bool {
	for i := 1; i < len(v); i++ {
		if v[i].Compare(v[i-1]) < 0 {
			return false
		}
	}

	return true
}
