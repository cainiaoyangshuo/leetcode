/**
 * @desc
 * @date 2022/6/13
 * @user yangshuo
 */

package BinaryTree

import (
	"fmt"
	"testing"
)

func Test_Btree(t *testing.T)  {
	nums := []int{2, 4, 6, 8, 10}

	root := CreateByBreadthFirstSearch(nums)

	level := levelOrder(root)
	fmt.Print(level)


	preRoot := CreateBTreeByPreOrder(nums)
	fmt.Print(preRoot)
	pre := levelOrder(preRoot)
	fmt.Print(pre)
}
