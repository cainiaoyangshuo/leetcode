<?php
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
	 * 我们可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，
	 * pp 指针和 qq 指针一开始都指向这棵树的根，随后 pp 右移时，qq 左移，pp 左移时，qq 右移。
	 * 每次检查当前 pp 和 qq 节点的值是否相等，如果相等再判断左右子树是否对称。
	 * @param TreeNode $root
	 * @return Boolean
	 */
	function isSymmetric($root) {
		return $this->isMirror($root, $root);
	}

	public function isMirror($p1, $p2)
	{
		if (empty($p1) && empty($p2)) return true;
		if (empty($p1) || empty($p2)) return false;

		return $p1->val===$p2->val && $this->isMirror($p1->left,$p2->right) && $this->isMirror($p1->right, $p2->left);
	}
}