package _01_对称二叉树

import BTree "a/BinaryTree"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * 递归
 */
func isSymmetric(root *BTree.TreeNode) bool {

	var isMirror func(p1, p2 *BTree.TreeNode) bool
	isMirror = func(p1, p2 *BTree.TreeNode) bool {

		if p1 == nil && p2 == nil {
			return true
		}

		if p1 == nil || p2 == nil {
			return false
		}
		return p1.Val == p2.Val && isMirror(p1.Left, p2.Right) && isMirror(p1.Right, p2.Left)
	}
	return isMirror(root, root)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 非递归，队列
 */
func isSymmetric_1(root *TreeNode) bool {
	queue := []*TreeNode{}
	j, k := root, root
	queue = append(queue, j)
	queue = append(queue, k)

	for len(queue) > 0 {
		j, k = queue[0], queue[1]
		queue = queue[2:]

		if j == nil && k == nil {
			continue
		}

		if (j == nil || k == nil) || j.Val != k.Val {
			return false
		}

		queue = append(queue, j.Left)
		queue = append(queue, k.Right)

		queue = append(queue, j.Right)
		queue = append(queue, k.Left)
	}

	return true
}