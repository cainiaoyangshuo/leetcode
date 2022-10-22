/**
 * @desc 226.翻转二叉树
 * @date 2022/6/21
 * @user yangshuo
 */

package _26_翻转二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 递归
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

/**
 * 迭代, 层序遍历，交换左右子树
 */
func invertTree_(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		curr.Left, curr.Right = curr.Right, curr.Left

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return root
}