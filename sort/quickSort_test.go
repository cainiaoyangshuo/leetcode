/**
 * @desc
 * @date 2022/9/2
 * @user yangshuo
 */

package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3, 1, 5, 2, 10, 8, 9, 22}
	res := QuickSort(nums)
	fmt.Printf("%+v\n", res)
}

func TestGeekQuickSort(t *testing.T) {
	nums := []int{3, 1, 5, 2, 10, 8, 9, 22}
	GeekQuickSort(nums)
	fmt.Print(nums)
}
