package calculate

func findIdxWithMax(arr []int, size int) (idxWithMax int) {
	for i := 1; i < size; i++ {
		if arr[i] > arr[idxWithMax] {
			idxWithMax = i
		}
	}
	return idxWithMax
}

func SelectionSort(arr []int) {
	N := len(arr)
	x := findIdxWithMax(arr, N)
	for j := N - 1; j > 0; j-- {
		arr[x], arr[j] = arr[j], arr[x]
		x = findIdxWithMax(arr, j)
	}
}
