package List

import "testing"

func TestGen(t *testing.T) {
	nums := []int{1,2,3,4,5}
	head := Gen(nums)
	PrintList(head)
}
