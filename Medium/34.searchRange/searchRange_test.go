package _4_searchRange

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T)  {
	nums := []int{}
	target := 8
	res := searchRange(nums, target)
	fmt.Printf("res : %+v\n", res)
}