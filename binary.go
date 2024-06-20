package sorting

func binarySearch[T Comparable[T]](v []T, e T) int {
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

	return i
}
