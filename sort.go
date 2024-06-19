// A collection of modified sorting algorithms for Go
package sorting

type Comparable[T any] interface {
	Compare(T) int
}

func IsSorted[T Comparable[T]](v []T) bool {
	for i := 1; i < len(v); i++ {
		if v[i].Compare(v[i-1]) < 0 {
			return false
		}
	}

	return true
}
