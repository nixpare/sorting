package sorting

// Rotate is an in-place not stable rotation algorithm:
//  + positive value for right rotations (from left to right, looping from the end back to the start)
//  + positive value for left rotations (from right to left, looping from the start to the end)
// The algorithm chooses the shortest path, so it may decide to invert the rotation.
// For example: if v has length 10 and it is called a right rotation of 6 places, the algorithm performs
// a left rotation of 4 places
func Rotate[E any](v []E, n int) {
	if n > 0 {
		if n < len(v) / 2 {
			rotateRight(v, n)
		} else {
			rotateLeft(v, len(v) - n)
		}
	} else if n < 0 {
		n = -n
		if n < len(v) / 2 {
			rotateLeft(v, n)
		} else {
			rotateRight(v, len(v) - n)
		}
	}
}

func rotateLeft[E any](v []E, n int) {
	start := 0
	x := v[start]
	
	for range n {
		for i := start - n + len(v) ; ; i -= n {
			v[i], x = x, v[i]

			if i == start {
				start++
				x = v[start]
				break
			}

			if i < n {
				start = i
				break
			}
		}
	}
}

func rotateRight[E any](v []E, n int) {
	start := 0
	x := v[start]
	
	for range n {
		for i := start + n ; ; i += n {
			v[i], x = x, v[i]

			if i + n == len(v) + start {
				v[start] = x
				start++
				x = v[start]
				break
			}

			if i + n >= len(v) {
				start = i + n - len(v)
				v[start], x = x, v[start]
				break
			}
		}
	}
}