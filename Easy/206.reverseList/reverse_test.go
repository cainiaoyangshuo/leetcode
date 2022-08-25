/**
 * @desc
 * @date 2022/5/21
 * @user yangshuo
 */

package _06_reverseList

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T)  {
	arr := []int{1,2,3,4,5}
	head := Gen(arr)
	PrintList(head)
	println()
	res := reverseList(head)
	PrintList(res)
}



func Gen(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	tail := head

	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{
			Val: nums[i],
		}
		tail = tail.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf(" %v", head.Val)
		head = head.Next
	}
	println()
}