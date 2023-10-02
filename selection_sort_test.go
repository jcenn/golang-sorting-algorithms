package main

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	for _, test := range tests {
		tmp := make([]int, len(test.have))
		copy(tmp, test.have)

		SelectionSort(tmp)
		for i := 0; i < len(tmp); i++ {
			if test.want[i] != tmp[i] {
				t.FailNow()
			}
		}
	}

}
