package sorting

func MergeSort[T Comparable[T]](v []T) {
	mergeSort(v, make([]T, len(v) / 4 + 1))
}

func mergeSort[T Comparable[T]](v []T, tmp []T) {
	switch len(v) {
	case 0, 1:
		return
	case 2:
		if v[0].Compare(v[1]) > 0 {
			v[0], v[1] = v[1], v[0]
		}
		return
	case 3:
		insertionSort(v)
		return
	}
	
	mid := len(v) / 2
	if mid % 2 == 1 {
		mid++
	}
	buffer := mid / 2

	mergeSort(v[mid:], tmp)        // Sort right part
	mergeSort(v[:buffer], tmp)     // Sort half of the left part

	left := merge(v, buffer, mid)  // Merges the left half part with the right full part to the right
	mergeSort(v[:left], tmp)       // Sort recursively the remaining left half part

	mergeParts(v, left, tmp)      // Merges the 1/4 to the left with the 3/4 to the right
}

func merge[T Comparable[T]](v []T, buffer int, right int) int {
	left, leftEnd := 0, buffer

	for left < leftEnd && right < len(v) {
		if v[left].Compare(v[right]) <= 0 {
			v[buffer], v[left] = v[left], v[buffer]
			left++
		} else {
			v[buffer], v[right] = v[right], v[buffer]
			right++
		}

		buffer++
	}

	for left < leftEnd {
		v[buffer], v[left] = v[left], v[buffer]
		left++; buffer++
	}

	for right < len(v) {
		v[buffer], v[right] = v[right], v[buffer]
		right++; buffer++
	}

	return left
}

func mergeParts[T Comparable[T]](v []T, mid int, tmp []T) {
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
		i++; k++
	}

	for j < len(v) {
		v[k] = v[j]
		j++; k++
	}
}