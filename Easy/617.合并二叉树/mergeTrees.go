/**
 * @desc
 * @date 2022/6/28
 * @user yangshuo
 */

package _17_合并二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 深度优先搜索+递归
 */
func mergeBTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeBTrees(root1.Left, root2.Left)
	root1.Right = mergeBTrees(root1.Right, root2.Right)

	return root1
}

/**
 * 广搜+队列
 */
func mergeBTrees_(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val+root2.Val}
	queue := []*TreeNode{root}
	queue1 := []*TreeNode{root1}
	queue2 := []*TreeNode{root2}

	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue = queue[1:]
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]

		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				queue = append(queue, left)
				queue1 = append(queue1, left1)
				queue2 = append(queue2, left2)
			} else if left1 != nil {
				node.Left = left1
			} else { // left2 != nil
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				queue = append(queue, right)
				queue1 = append(queue1, right1)
				queue2 = append(queue2, right2)
			} else if right1 != nil {
				node.Right = right1
			} else { // right2 != nil
				node.Right = right2
			}
		}
	}

	return root
}