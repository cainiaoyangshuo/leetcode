package _43_diameterOfBinaryTree

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	if root == nil {
		return maxDiameter
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lh := dfs(root.Left)
		rh := dfs(root.Right)

		maxDiameter = max(maxDiameter, lh + rh)
		return 1 + max(lh, rh)
	}

	dfs(root)

	return maxDiameter
}

func max(a,b int ) int {
	if a > b {
		return a
	}
	return b
}