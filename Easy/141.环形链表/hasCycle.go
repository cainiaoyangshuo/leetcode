/**
 * @desc
 * @date 2025/5/22
 * @user yangshuo
 */

package _41_环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 快慢指针不指向同一节点，不利于判断重复
	fast, slow := head.Next, head

	for fast != slow {
		// fast走到头则证明无环，因为快指针走两步，未避免空指针，fast和fast.Next都要校验
		if fast == nil || fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
