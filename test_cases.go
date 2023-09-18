package main

var tests []Test = []Test{
	{
		have: []int{5, 2, 1, -2, 3, 12, 0, -55},
		want: []int{-55, -2, 0, 1, 2, 3, 5, 12},
	},
}
