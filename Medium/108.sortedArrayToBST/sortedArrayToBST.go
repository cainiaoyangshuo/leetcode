/**
 * @desc 108.升序数组转二叉搜索树
 * @date 2022/7/12
 * @user yangshuo
 */

package _08_sortedArrayToBST

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 中序遍历，中点作为树的根节点
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var inorder func(nums []int, left, right int) *TreeNode
	inorder = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = inorder(nums, left, mid - 1)
		root.Right = inorder(nums, mid + 1, right)
		return root
	}

	return inorder(nums, 0, len(nums)-1)
}
