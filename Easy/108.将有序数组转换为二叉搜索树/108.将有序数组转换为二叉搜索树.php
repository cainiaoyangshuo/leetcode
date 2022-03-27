<?php
/**
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	示例:

	给定有序数组: [-10,-3,0,5,9],

	一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

	 0
	/ \
  -3   9
  /   /
-10  5

 * @date 2020/9/1
 */

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution
{
    /**
     * dfs 二叉搜索树中序遍历是升序的，有序数组任意元素作为根节点进行中序遍历，
     * 以该元素左边部分作为左子树，右边作为右子树，为高度平衡，选择数组中间元素。
     * @param Integer[] $nums
     * @return TreeNode
     */
    function sortedArrayToBST($nums)
    {
        return $this->dfs($nums, 0, count($nums) - 1);
    }

    function dfs($nums, $low, $high)
    {
        if ($low > $high) {
            return null;
        }

        // 以数组中间作为root节点
        $mid = $low + round(($high - $low) / 2);
        $root = new TreeNode($nums[$mid]);

        // 递归构建root的左子树与右子树
        $root->left = $this->dfs($nums, $low, $mid - 1);
        $root->right = $this->dfs($nums, $mid + 1, $high);

        return $root;
    }
}
