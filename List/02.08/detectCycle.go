package _2_08

// ListNode /**
type ListNode struct {
	 Val int
	 Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil || slow == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		//判断是否有环
		if fast == slow {
			break
		}
	}

	start := head
	for {
		if start == slow {
			return start
		}
		start = start.Next
		slow = slow.Next
	}
}
