<?php
/**
 * @date 2020/8/26
 */

class Solution {

	/**
	 * @param TreeNode $root
	 * @return Integer[]
	 */
	function inorderTraversal($root)
	{
		if (empty($root)) {
			return [];
		}

		$res = [];
		$stack = [];
		while (!empty($root) || !empty($stack)) {
			if (!empty($root)) {
				array_push($stack, $root);
				$root = $root->left;
			} else {
				$root = array_pop($stack);
				$res[] = $root->val;
				$root = $root->right;
			}
		}

		return $res;

	}
}