package main

// Sorts given slice by keeping the sorted elements at the beginning 
// and inserting new elements onto correct positions as it iterates over them
// Time complexity O(n^2)
func insertion_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			swap(j, j-1, arr)
			j -= 1
		}
	}
}
