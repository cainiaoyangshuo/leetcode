package _2_trap

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T)  {
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	ans := trap(nums)
	fmt.Println(ans)

	nums2 := []int{4,2,0,3,2,5}
	ans = trap(nums2)
	println(ans)
}
