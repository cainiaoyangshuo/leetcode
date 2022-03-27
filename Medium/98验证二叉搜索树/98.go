package _98验证二叉搜索树

import "math"

/**
 https://leetcode-cn.com/problems/validate-binary-search-tree/
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	var inOrder func(root *TreeNode) bool
	inOrder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !inOrder(root.Left) {
			return false
		}

		if root.Val > pre {
			pre = root.Val
		} else {
			return false
		}

		return inOrder(root.Right)
	}
	return inOrder(root)
}

//func inOrder(root *TreeNode, a int) bool {
//	if root != nil {
//		if !inOrder(root.Left, a) {
//			return false
//		}
//
//		if root.Val > a {
//			a = root.Val
//		} else {
//			return false
//		}
//
//		if !inOrder(root.Right, a) {
//			return false
//		}
//	}
//
//	return true
//}