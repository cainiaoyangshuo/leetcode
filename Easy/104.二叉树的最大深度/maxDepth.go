/**
 * @desc
 * @date 2022/6/21
 * @user yangshuo
 */

package _04_二叉树的最大深度

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	queue := []*TreeNode{root}

	for queue != nil {
		tmp := queue
		queue = nil

		depth++

		for _, node := range tmp {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}
