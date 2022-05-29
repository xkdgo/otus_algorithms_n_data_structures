package calculate

func BubbleSort(arr []int) {
	N := len(arr)
	for i := 1; i <= N-1; i++ {
		for j := 0; j < N-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
