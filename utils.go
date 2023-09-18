package main

type Test struct {
	have []int
	want []int
}

func swap(i1 int, i2 int, arr []int) {
	tmp := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = tmp
}
