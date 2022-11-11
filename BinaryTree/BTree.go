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

//层序创建二叉树
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

//前序遍历创建二叉树
func CreateBTreeByPreOrder(input []int)  *TreeNode {
	if input == nil {
		return nil
	}

	var pre func(input []int, i int) *TreeNode
	pre = func(input []int, i int) *TreeNode {
		if i >= len(input) {
			return nil
		}

		root := CreateNode(input[i])
		i++
		root.Left = pre(input, i)
		root.Right = pre(input, i)
		return root
	}

	root := pre(input, 0)
	return root
}

/**
 * 二叉树层序遍历
 */
func levelOrder(root *TreeNode) (res [][]int) {

	queue := []*TreeNode{
		root,
	}

	for queue != nil {
		tmp := queue
		queue = nil

		level := []int{}
		for _, node := range tmp {
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return
}

