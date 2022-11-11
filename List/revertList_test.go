/**
 * @desc
 * @date 2022/11/2
 * @user yangshuo
 */

package List

import "testing"

func TestRevertList(t *testing.T) {
	nums := []int{1,5,6,10,20}
	head := Gen(nums)

	head = RevertList(head)
	PrintList(head)
}
