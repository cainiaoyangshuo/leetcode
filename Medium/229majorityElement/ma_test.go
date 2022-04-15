package _29majorityElement

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2,1,1,3,1,4,5,6}
	res := MajorityElement(nums)
	fmt.Printf("%d\n", res)
}
