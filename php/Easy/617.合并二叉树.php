<?php
/**
 * @befie            #功能说明
 * @version          #版本
 * @author   yangshuo
 * @date     2020/9/23
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
class Solution {

	/**
	 * @param TreeNode $t1
	 * @param TreeNode $t2
	 * @return TreeNode
	 */
	function mergeTrees($t1, $t2) {
		if (empty($t1)) {
			return $t2;
		}

		if (empty($t2)) {
			return $t1;
		}

		$merge = new TreeNode($t1->val + $t2->val);
		$merge->left = $this->mergeTrees($t1->left, $t2->left);
		$merge->right = $this->mergeTrees($t1->right, $t2->right);

		return $merge;
	}
}