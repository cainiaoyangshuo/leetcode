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
