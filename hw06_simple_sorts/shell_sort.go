package calculate

import "math"

func ShellSortClassical(arr []int) {
	N := len(arr)
	for gap := N / 2; gap > 0; gap /= 2 {
		for i := gap; i < N; i++ {
			tmp := arr[i]
			j := i
			for ; j >= gap; j -= gap {
				if tmp < arr[j-gap] {
					arr[j] = arr[j-gap]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
}

func ShellSort2Kminus1(arr []int) {
	N := len(arr)
	K := int(math.Ceil(math.Log(float64(N)) / math.Log(float64(2))))
	for gap := int(math.Pow(2, float64(K))) - 1; gap > 0; gap /= 2 {
		for i := gap; i < N; i++ {
			tmp := arr[i]
			j := i
			for ; j >= gap; j -= gap {
				if tmp < arr[j-gap] {
					arr[j] = arr[j-gap]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
}

func ShellSort3Kminus1(arr []int) {
	N := len(arr)
	K := int(math.Ceil(math.Log(float64(N)) / math.Log(float64(2))))
	for gap := (int(math.Pow(3, float64(K))) - 1); gap > 0; gap /= 2 {
		for i := gap; i < N; i++ {
			tmp := arr[i]
			j := i
			for ; j >= gap; j -= gap {
				if tmp < arr[j-gap] {
					arr[j] = arr[j-gap]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
}
