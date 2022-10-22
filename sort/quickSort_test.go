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
	nums := []int{5, 2, 3, 9}
	res := QuickSort(nums)
	fmt.Printf("%+v\n", res)
}


