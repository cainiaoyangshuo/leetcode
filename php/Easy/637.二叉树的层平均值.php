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
	 * @param TreeNode $root
	 * @return Float[]
	 */
	function averageOfLevels($root) {
		if (empty($root)) {
			return null;
		}

		$queue = [];
		$res = [];
		$re = [];
		$queue[] = $root;
		$level = 0;
		while ($count = count($queue)) {
			for ($i = 0; $i < $count; $i++) {
				$curr = array_shift($queue);
				$res[$level][] = $curr->val;

				if ($curr->left) {
					$queue[] = $curr->left;
				}

				if ($curr->right) {
					$queue[] = $curr->right;
				}
			}
			$curLevel = $res[$level];
			$re[] = array_sum($curLevel) / count($curLevel);
			$level++;
		}

		return $re;
	}
}