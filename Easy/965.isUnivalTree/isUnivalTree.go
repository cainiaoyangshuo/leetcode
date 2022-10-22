/**
 * @desc 965.单值二叉树 https://leetcode.cn/problems/univalued-binary-tree/
 * @date 2022/5/25
 * @user yangshuo
 */

package _65_isUnivalTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 中序遍历，递归
 */
func isUnivalTree(root *TreeNode) bool {

	first := root.Val
	res := true

	var innor func(node *TreeNode)
	innor = func(node *TreeNode) {
		if node == nil {
			return
		}
		innor(node.Left)
		if node.Val != first {
			res = false
			return
		}
		first = root.Val
		innor(node.Right)
	}

	innor(root)

	return res
}


/**
 * 迭代
 */
func isUnivalTree1(root *TreeNode) bool {
	stack := []*TreeNode{}

	first := root.Val

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Val != first {
			return false
		}
		first = root.Val

		root = root.Right
	}

	return true
}