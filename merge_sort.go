package main

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[0:mid]
	right := arr[mid:]

	return merge(MergeSort(left), mergeSort(right))
}

func merge(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0
	k := 0
	result := make([]int, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i += 1
		} else {
			result[k] = arr2[j]
			j += 1
		}
		k += 1
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i += 1
		k += 1
	}

	for j < len(arr2) {
		result[k] = arr2[j]
		j += 1
		k += 1
	}

	return result
}
