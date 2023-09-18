package main

// Sorts given slice by iterating through all elements and comparing neighbors,
// swapping them if they're not in the right order. Algorithm will iterate until array is sorted,
// omitting last elements as they're guaranteed to be sorted
func bubble_sort(arr []int) {
	j := 0
	for j < len(arr)-1 {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i+1] < arr[i] {
				swap(i, i+1, arr)
			}
		}
		j += 1
	}
}
