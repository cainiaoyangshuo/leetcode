package _40_SearchMatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	nums := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	target := 5
	res := SearchMatrix(nums, target)
	println(res)
}
