package sorting

// Rotate is an in-place rotation algorithm:
//  + positive value for right rotations
//  + negative value for left rotation
func Rotate[T any](v []T, n int) {
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

func rotateLeft[T any](v []T, n int) {
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

func rotateRight[T any](v []T, n int) {
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