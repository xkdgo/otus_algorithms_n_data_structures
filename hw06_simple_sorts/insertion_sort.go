package calculate

func InsertionSortSwapEach(arr []int) {
	N := len(arr)
	for i := 1; i <= N-1; i++ {
		j := i - 1
		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
}

func InsertionSortOneHit(arr []int) {
	N := len(arr)
	for i := 0; i <= N-1; i++ {
		x := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}

func InsertionSortWithBinarySearch(arr []int) {
	N := len(arr)
	for i := 1; i < N; i++ {
		j := i - 1
		x := arr[i]
		// find location where selected should be inseretd
		loc := binarySearch(arr, x, 0, j)
		// Move all elements after location to create space
		for j >= loc {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}

func binarySearch(arr []int, val int, startIdx, endIdx int) (positionIdx int) {
	if startIdx == endIdx {
		if arr[startIdx] > val {
			return startIdx
		} else {
			return startIdx + 1
		}
	}
	if startIdx > endIdx {
		return startIdx
	}
	mid := (startIdx + endIdx) / 2
	switch {
	case arr[mid] < val:
		return binarySearch(arr, val, mid+1, endIdx)
	case arr[mid] > val:
		return binarySearch(arr, val, startIdx, mid-1)
	default:
	}
	return mid + 1
}
