/**
 * @desc 83.删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
 * @date 2022/5/25
 * @user yangshuo
 */

package _3_deleteDuplicates

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * 定义指针curr指向head，如果curr=curr.Next，则将curr.Next移除，否则curr = curr.Next
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head

	for curr.Next != nil {
		if curr.Next.Val == curr.Val {
			//移除curr.Next，即将next的next指向curr.Next
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
