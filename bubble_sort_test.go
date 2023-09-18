package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, test := range tests{
		tmp :=  make([]int, len(test.have))
		copy(tmp, test.have)

		bubble_sort(tmp)
		for i:=0; i < len(tmp); i++ {
			if test.want[i] != tmp[i]{
				t.FailNow()
			}
		}
	}
}
