/**
 * @desc
 * @date 2022/5/21
 * @user yangshuo
 */

package _06_reverseList

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * 1.辅助栈迭代 O(n), O(n)
 */
func reverseList(head *ListNode) *ListNode {
	stack := []int{}
	tail := head
	var res,t *ListNode

	for tail != nil {
		stack = append(stack, tail.Val)
		tail = tail.Next
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if res == nil {
			res = &ListNode{Val:top}
			t = res
		} else {
			t.Next = &ListNode{Val:top}
			t = t.Next
		}
		stack = stack[:len(stack)-1]
	}

	return res
}

/**
 * 2.不借助辅助栈迭代 o(n) O(1) 借助一个空指针
 */
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	//辅助指针
	curr := head

	for curr != nil {
		//临时记录curr.Next，防止curr.Next指向pre时与后面的结点断开
		next := curr.Next
		//把curr.next指向pre
		curr.Next = pre

		//迭代下个结点
		pre = curr
		curr = next
	}

	return pre
}

/**
 * 3.递归
 */
func reverseList2(head *ListNode) *ListNode {
	//终止条件
	if head == nil || head.Next == nil {
		return nil
	}

	//
	p := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}