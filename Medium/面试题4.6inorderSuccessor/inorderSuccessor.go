/**
 * @desc
 * @date 2022/5/16
 * @user yangshuo
 */

package 面试题4_6inorderSuccessor

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 1.中序遍历 O(N),O(N)
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	cur := root
	var pre *TreeNode = nil
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {

		//根节点左子树入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		//从最左边的左子树开始（最小数）
		cur = stack[len(stack)-1]
		//出栈
		stack = stack[:len(stack)-1]

		//pre
		if pre == p {
			return cur
		}

		//pre记为上一个节点
		pre = cur
		cur = cur.Right
	}

	return nil
}

/**
 * 2.二分 O(N) O(1)
 */
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var ans *TreeNode
	cur := root

	//二分
	for cur != nil {
		if cur.Val > p.Val {
			ans = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return ans
}

