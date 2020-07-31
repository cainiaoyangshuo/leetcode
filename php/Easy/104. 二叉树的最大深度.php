<?php
/**
 * @date 2020/7/15
 */

/**
 * 给定一个二叉树，找出其最大深度。

	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	说明: 叶子节点是指没有子节点的节点。

	示例：
	给定二叉树 [3,9,20,null,null,15,7]，

	3
	/ \
	9  20
	/  \
	15   7
	返回它的最大深度 3 。

 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

	public function __call( $name, $arguments ) {
		// TODO: Implement __call() method.
		$ref = new ReflectionClass(__CLASS__);
		$funcs = $ref->getMethods();
		$func = $funcs[count($funcs)-1]->name;
		return $this->$func($arguments);
	}

	/**
	 * @param TreeNode $root
	 * @return Integer
	 */
	function maxDepth($root) {
		if (empty($root)) {
			return 0;
		}


		$left = $this->maxDepth($root->left);
		$right = $this->maxDepth($root->right);

		return max($left, $right) + 1;
	}
}

$obj = new Solution();
$n = [3,9,20,null,null,15,7];
$res = $obj->test($n);
var_dump($res);