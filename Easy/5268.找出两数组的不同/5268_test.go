package _268_找出两数组的不同

import (
	"fmt"
	"testing"
)

func Test_findDifference(t *testing.T)  {
	nums1 := []int{1,2,3,3}
	nums2 := []int{1,1,2,2}

	res := findDifference(nums1, nums2)
	fmt.Printf("%v\n ", res)
}
