// A collection of modified sorting algorithms for Go
package sorting

type Comparable[T any] interface {
	Compare(T) int
}
