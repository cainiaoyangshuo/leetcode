package _69majorityElement

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2,2,1,1,1,2,1,1,2}
	res := MajorityElement(nums)
	fmt.Printf("%d\n", res)
}
