package _29levelOrder
/**
 * @desc N叉数层序遍历 https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
 * @date 2022/4/18
 * @user yangshuo
 */


type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*Node{root}
	for queue != nil {
		// 每层
		level := []int{}
		temp := queue // temp用来遍历当前队列
		queue = nil // 清空队列，便于记录下一层
		for _, node := range temp {
			level = append(level, node.Val)
			// 将所有子节点放入队列
			queue = append(queue, node.Children...)
		}

		res = append(res, level)
	}

	return res
}