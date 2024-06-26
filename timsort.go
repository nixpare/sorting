package sorting

var (
	TimSortRun = 32
)

func TimSort[T Comparable[T]](v []T) {
	// Sort individual subarrays of size RUN
	for i := 0; i < len(v); i += TimSortRun {
		from, to := i, min(i+TimSortRun, len(v))
		InsertionSort(v[from:to])
	}

	tmp := newBuffer(v, 0)

	// Start merging from size RUN (or 32).
	// It will merge
	// to form size 64, then 128, 256
	// and so on ....
	for size := TimSortRun; size < len(v); size *= 2 {

		// pick starting point of
		// left sub array. We
		// are going to merge
		// arr[left..left+size-1]
		// and arr[left+size, left+2*size-1]
		// After every merge, we
		// increase left by 2*size
		for left := 0; left < len(v); left += 2 * size {

			// Find ending point of
			// left sub array
			// mid+1 is starting point
			// of right sub array
			mid := left + size
			right := min(mid+size, len(v))

			// merge sub array arr[left.....mid] &
			// arr[mid+1....right]
			if mid < right {
				mergeExternalBuffer(v[left:right], size, tmp)
			}
		}
	}
}
