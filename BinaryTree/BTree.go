package BinaryTree

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateNode(item int) *TreeNode {

	if item == -1 {
		return nil
	}

	return &TreeNode{item, nil, nil}
}

func CreateByBreadthFirstSearch(inputList []int) *TreeNode {

	queue := []*TreeNode{}
	root := CreateNode(inputList[0])
	queue = append(queue, root)

	inputList = inputList[1:]

	for len(inputList) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		curNode.Left = CreateNode(inputList[0])
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		inputList = inputList[1:]

		curNode.Right = CreateNode(inputList[0])
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
		inputList = inputList[1:]

	}

	return root
}

func Inorder(root *TreeNode) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Printf("%v\n", root.Val)
	Inorder(root.Right)
}

/**
 * 二叉树层序遍历
 */
func levelOrder(root *TreeNode) {

}

