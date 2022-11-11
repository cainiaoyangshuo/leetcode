/**
 * @desc
 * @date 2022/11/7
 * @user yangshuo
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3, 1, 5, 2, 10, 8, 9, 22}

	res := QuickSort(nums)
	fmt.Println(res)
}
