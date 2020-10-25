package mergesort

// Sort sorts the array ascending with the complexity of O(n logn).
func Sort(arr []int) []int {
	l := len(arr)

	// skip
	if l < 2 {
		return arr
	}

	// create a copy
	arrCopy := make([]int, l)
	for i := 0; i < l; i++ {
		arrCopy[i] = arr[i]
	}

	// find mid
	mid := l / 2

	// divide and sort
	leftArr := Sort(arrCopy[:mid])
	rightArr := Sort(arrCopy[mid:])

	// merge
	leftLen, rightLen := len(leftArr), len(rightArr)
	leftInd, rightInd, arrInd := 0, 0, 0

	for leftInd < leftLen && rightInd < rightLen {

		lVal, rVal := leftArr[leftInd], rightArr[rightInd]

		// compare
		if lVal < rVal {
			arr[arrInd] = lVal
			leftInd++
		} else {
			arr[arrInd] = rVal
			rightInd++
		}

		arrInd++
	}

	// fill with the other
	for leftInd < leftLen {
		arr[arrInd] = leftArr[leftInd]
		leftInd++
		arrInd++
	}
	for rightInd < rightLen {
		arr[arrInd] = rightArr[rightInd]
		rightInd++
		arrInd++
	}

	return arr
}
