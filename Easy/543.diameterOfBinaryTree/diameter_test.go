package _43_diameterOfBinaryTree

import (
	"fmt"
	"testing"
)

func TestDia(t *testing.T)  {
	nums := []int{1,2,4,5,6,8,9}
	node := CreateByBreadthFirstSearch(nums)

	max := diameterOfBinaryTree(node)
	fmt.Printf("max: %v\n", max)
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