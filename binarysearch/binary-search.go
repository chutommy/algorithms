package binarysearch

// Search searches for an element K in the given array and returns
// its index. If an element is not found, (-1) is returned.
// Given array must be sorted otherwise the algorithm would
// not work properly. Algorithm's complexity is of O(logn).
func Search(K int, arr []int) int {
	l := len(arr)

	// declare start/end
	a, b := 0, l-1

	// search
	for a <= b {

		// find middle index
		mid := (a + b) / 2

		if arr[mid] == K {
			return mid // found
		}

		// divide
		if arr[mid] > K {
			b = mid - 1
		} else {
			a = mid + 1
		}
	}

	// K not found
	return -1
}
