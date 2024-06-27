package sorting

// BinarySearch implements the standard search algorithm over a sorted slice.
// Returns -1 if the element was not found
func BinarySearch[S ~[]E, E Ordered[E]](v S, e E) int {
	from, to, i := 0, len(v), len(v)/2

	for i >= from && i < to {
		compare := v[i].Compare(e)
		if compare == 0 {
			for i-1 >= from && v[i-1].Compare(e) == 0 {
				i--
			}

			return i
		}

		if compare > 0 {
			to = i
		} else {
			from = i + 1
		}

		i = (to + from) / 2
	}

	if from == to {
		return -1
	}
	return i
}
