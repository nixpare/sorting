package sorting

// MergeExternal merges two parts of a slice with the help of an external
// buffer. This algorithm is stable and not in-place (but does not allocate itself
// anuthing).
// The caller must assure that:
//  - mid is a valid index for v
//  - tmp is at least as long as the shorted part of v
// This algorithm is based on the standard merge from mergesort but has a few optimizations:
//  - detects if the two parts are already merged, resulting in a nop
//  - detects if the two parts are swapped, reducing the number of comparisons
//  - choses to merge from left to right or vice versa in order to reduce copies to the buffer area
func MergeExternal[S ~[]E, E Ordered[E]](v S, mid int, tmp S) {
	if v[mid-1].Compare(v[mid]) < 0 {
		return
	}

	if v[0].Compare(v[len(v)-1]) > 0 {
		Swap(v, mid, tmp)
		return
	}

	if mid <= len(v)-mid {
		mergeExternalBufferFromLeft(v, mid, tmp)
	} else {
		mergeExternalBufferFromRight(v, mid, tmp)
	}
}

func mergeExternalBufferFromLeft[S ~[]E, E Ordered[E]](v S, mid int, tmp S) {
	copy(tmp, v[:mid])

	i, j, k := 0, mid, 0
	for i < mid && j < len(v) {
		if tmp[i].Compare(v[j]) <= 0 {
			v[k] = tmp[i]
			i++
		} else {
			v[k] = v[j]
			j++
		}

		k++
	}

	for i < mid {
		v[k] = tmp[i]
		i++
		k++
	}

	for j < len(v) {
		v[k] = v[j]
		j++
		k++
	}
}

func mergeExternalBufferFromRight[S ~[]E, E Ordered[E]](v S, mid int, tmp S) {
	n := copy(tmp, v[mid:])
	tmp = tmp[:n]

	i, j, k := mid-1, len(tmp)-1, len(v)-1
	for i >= 0 && j >= 0 {
		if v[i].Compare(tmp[j]) > 0 {
			v[k] = v[i]
			i--
		} else {
			v[k] = tmp[j]
			j--
		}

		k--
	}

	for i >= 0 {
		v[k] = v[i]
		i--
		k--
	}

	for j >= 0 {
		v[k] = tmp[j]
		j--
		k--
	}
}

// MergeInternal merges two parts of a slice with the help of a buffer area.
// This algorithm is in-place and not stable
// The input slice must be in this state:
//  - v[ : buffer] is the fist sorted part to be merged
//  - v[ buffer : right ] is the buffer area and must have the same size as the first part
//  - v[ right : ] is the second sorted part to be merged and must have at least the size of the buffer
// The result will be a slice where:
//  - v[ : buffer ] contains the elements before in the buffer area (unsorted)
//  - v[ buffer : ] contains the merged sorted result
// It is optimal to use this function if the slice is made like this:
//      v:  { [  1 / 4  ]  [  1 / 4  ]  [  1 / 2  ] }
//       0 -^      buffer -^     right -^
func MergeInternal[S ~[]E, E Ordered[E]](v S, buffer int, right int) {
	left := 0

	for buffer < right && right < len(v) {
		if v[left].Compare(v[right]) <= 0 {
			v[buffer], v[left] = v[left], v[buffer]
			left++
		} else {
			v[buffer], v[right] = v[right], v[buffer]
			right++
		}

		buffer++
	}

	for buffer < right {
		v[buffer], v[left] = v[left], v[buffer]
		left++
		buffer++
	}
}

// Swap swaps the two parts of a slice. This procedure is not in-place
// and not stable.
// The caller must assure that the length of tmp is at least equal to
// the shortest part of the slice.
// For and in-place alternative, see the Rotate procedure
func Swap[S ~[]E, E Ordered[E]](v S, mid int, tmp S) {
	if mid <= len(v)-mid {
		n := copy(tmp, v[:mid])

		i := 0
		for j := mid; j < len(v); j++ {
			v[i], v[j] = v[j], v[i]
			i++
		}

		copy(v[i:], tmp[:n])
	} else {
		n := copy(tmp, v[mid:])

		i := len(v) - 1
		for j := mid - 1; j >= 0; j-- {
			v[i], v[j] = v[j], v[i]
			i--
		}

		copy(v[:i], tmp[:n])
	}
}
