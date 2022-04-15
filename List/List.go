package List

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
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
		fmt.Printf(" %v\n", head.Val)
		head = head.Next
	}
}