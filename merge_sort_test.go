package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		tmp := make([]int, len(test.have))
		copy(tmp, test.have)

		tmp = MergeSort(tmp)

		s := "result: "

		for i := 0; i < len(tmp); i++ {
			s += fmt.Sprintf("%d, ", tmp[i])
		}
		fmt.Println(s)
		for i := 0; i < len(tmp); i++ {
			if test.want[i] != tmp[i] {
				t.FailNow()
			}
		}
	}

}
