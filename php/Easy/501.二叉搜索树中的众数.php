<?php
/**
 * @befie            #功能说明
 * @version          #版本
 * @author   yangshuo
 * @date     2020/9/24
 */

/**
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

	假定 BST 有如下定义：

	结点左子树中所含结点的值小于等于当前结点的值
	结点右子树中所含结点的值大于等于当前结点的值
	左子树和右子树都是二叉搜索树
	例如：
	给定 BST [1,null,2,2],
	
	1
     \
     2
	/
   2
	返回[2].

	提示：如果众数超过1个，不需考虑输出顺序

	进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）


 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

	/**
	 * @param TreeNode $root
	 * @return Integer[]
	 */
	function findMode($root) {
		if (empty($root)) {
			return [];
		}

		$stack = [];
		$count = 0;
		$maxCount = 1;
		$base = null;
		$res = [];

		while (!empty($stack) || !empty($root)) {
			if (!empty($root)) {
				$stack[] = $root;
				$root = $root->left;
			} else {
				$root = array_pop($stack);
				$val = $root->val;
				if ($val == $base) {
					++$count;
				} else {
					$count = 1;
					$base = $val;
				}
				if ($count == $maxCount) {
					$res[] = $val;
				}

				if ($count > $maxCount) {
					$maxCount = $count;
					$res = null;
					$res[] = $val;
				}
				$root = $root->right;
			}
		}

		return $res;
	}
}