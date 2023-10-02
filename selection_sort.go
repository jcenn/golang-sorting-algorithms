package main

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			swap(i, minIndex, arr)
		}
	}
}
