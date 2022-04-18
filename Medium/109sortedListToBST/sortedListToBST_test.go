package _09sortedListToBST

import (
	BTree "a/BinaryTree"
	"a/List"
	"testing"
)

/**
 * @desc
 * @date 2022/4/17
 * @user yangshuo
 */

func Test_sortedListToBST(t *testing.T)  {
	arr := []int{-10,-3,0,5,9}
	// 生成链表
	head := List.Gen(arr)
	List.PrintList(head)

	// 生成二叉树
	root := sortedListToBST(head)

	// 中序遍历二叉树
	println("")
	BTree.Inorder(root)

	// 前序遍历二叉树

}