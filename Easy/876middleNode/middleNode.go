/**
 * @desc
 * @date 2022/4/27
 * @user yangshuo
 */
package _76middleNode


  type ListNode struct {
      Val int
      Next *ListNode
  }
func middleNode(head *ListNode) *ListNode {
	left, right := head, head

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}

	return left
}
