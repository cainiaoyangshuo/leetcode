/**
 * @desc 2.两数相加 https://leetcode.cn/problems/add-two-numbers/
 * @date 2022/5/13
 * @user yangshuo
 */

package __两数相加


type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * 思路：模拟，因为是逆序存储，进位加到下一个结点，直接遍历相加即可。如果遍历完进位不为0，还需要再新增结点
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//进位
	carry := 0
	//head为输出结点，tail为辅助结点
	var head, tail *ListNode

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0

		//遍历
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		//相加
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		//初始化输出链表
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
