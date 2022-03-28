package _60

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T)  {
	nums := []int{1,1,1}
	k := 2

	count := subarraySum(nums, k)
	fmt.Printf("%v\n", count)
}
