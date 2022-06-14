package calculate

func HeapSort(arr []int) {
	N := len(arr)
	for root := ((N - 1) - 1) / 2; root >= 0; root-- {
		heapify(arr, root, N)
	}
	for j := N - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		heapify(arr, 0, j)
	}
}

func heapify(arr []int, root, size int) {
	leftChildIdx := 2*root + 1
	rightChildIdx := leftChildIdx + 1
	indexWithMax := root
	if leftChildIdx < size && arr[leftChildIdx] > arr[indexWithMax] {
		indexWithMax = leftChildIdx
	}
	if rightChildIdx < size && arr[rightChildIdx] > arr[indexWithMax] {
		indexWithMax = rightChildIdx
	}
	if indexWithMax == root {
		return
	}
	arr[indexWithMax], arr[root] = arr[root], arr[indexWithMax]
	heapify(arr, indexWithMax, size)
}
