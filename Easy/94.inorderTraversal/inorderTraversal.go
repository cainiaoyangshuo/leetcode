/**
 * @desc 94.二叉树中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/
 * @date 2022/5/24
 * @user yangshuo
 */

package _4_inorderTraversal


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 递归
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return res
}

/**
 * 迭代
 */
func inorderTraversal1(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}

	for root != nil || len(stack) > 0 {
		//左子树入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//取栈顶，并出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, root.Val)

		//处理右节点
		root = root.Right
	}

	return res
}

/**
 * morris 中序遍历 将非递归遍历空间复杂度降为O(1), 设当前节点为x
 * 1.若x无左孩子，则先将x加入答案数组，再访问x右孩子，即x=x.right
 * 2.若有左孩子，则找到x左子树上最右的节点，（即左子树中序遍历的最后一个节点，x 在中序遍历中的前驱节点），我们记为 predecessor，根据predecessor右孩子是否为空做如下操作：
 * ①.若 predecessor右孩子为空， 则将其右孩子指向 x, 即 x = x.left
 * ② 若 predecessor 右孩子不为空 则此时其右孩子指向 x, 说明我们已经遍历完 x 的左子树，将 predecessor 右孩子置为空，将x的值加入答案数组，然后访问 x 的右孩子， 即 x = x.right
 */
func inorderTraversal2(root *TreeNode) (res []int) {

	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}