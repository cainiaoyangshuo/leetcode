package _09_有序链表转换二叉搜索树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
	分治
	找到链表中心点，以中心点左边的节点构建左子树，以中心右边构建右子树
	链表中心点可通过快慢指针找到
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	slow, fast := head, head
	prev := &ListNode{}
	// 因为快指针要走两步，所以要同时判断fast及fast的下一个是否为空
	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	//此时slow即为链表中点
	prev.Next = nil

	//递归构建bst
	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}