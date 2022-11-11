/**
 * @desc
 * @date 2022/11/2
 * @user yangshuo
 */

package List

func RevertList(head *ListNode) *ListNode {

	curr := head
	var pre *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = pre

		pre = curr
		curr = next
	}

	return pre
}
