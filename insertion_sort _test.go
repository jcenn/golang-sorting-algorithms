package main

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{-1, 5, 6, 2, 4, 3, 2, 1}

	fmt.Printf("before sorting: %+v \n", arr)
	insertion_sort(arr)
	fmt.Printf("after sorting: %+v \n", arr)

}
