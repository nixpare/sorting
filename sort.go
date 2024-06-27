// A collection of modified sorting algorithms for Go
package sorting

type Ordered[E any] interface {
	Compare(E) int
}

func IsSorted[S ~[]E, E Ordered[E]](v S) bool {
	for i := 1; i < len(v); i++ {
		if v[i].Compare(v[i-1]) < 0 {
			return false
		}
	}

	return true
}
