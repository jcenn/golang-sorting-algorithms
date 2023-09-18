package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// tests := []Test{
	// 	{have: []int{5, 2 ,1 ,-2, 3,12, 0, -55}, want: []int{}},
	// }
	arr := []int{5, 2 ,1 ,-2, 3,12, 0, -55}

	fmt.Printf("before sorting: %+v \n", arr)
	bubble_sort(arr)
	fmt.Printf("after sorting: %+v \n", arr)
	// tests := []struct {
	// 	number int
	// 	want   int
	// }{
	// 	{number: 10, want: 100},
	// 	{number: 5, want: 25},
	// 	{number: 7, want: 49},
	// }
	// for _, tc := range tests {
	// 	if tc.want != 0 {
	// 		t.Fatal("oh no")
	// 		// t.Fatalf("Test %d failed — Expected %d, got %d", i+1, tc.want, result)
	// 	}
	// }

}
