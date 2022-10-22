/**
 * @desc 19.删除链表的倒数第n个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 * @date 2022/5/26
 * @user yangshuo
 */

package _9_removeNthFromEnd

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * 双指针，p1, p2, p1先走n步，然后p2再开始走，当p1走到末尾，p2刚好走到倒数第n个结点。
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	res := &ListNode{0, head}
	p1, p2 := head, res

	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return res.Next
}
