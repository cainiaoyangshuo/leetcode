<?php
class TreeNode
{
	public $left = null;
	public $right = null;
	public $value = null;
}

function reverseTree($root)
{
	if ($root === null) {
		return null;
	}

	$left = $root->left;
	$right = $root->right;
	$root->left = $this->reverseTree($right);
	$root->right = $this->reverseTree($left);
	return $root;

}

/**
 * @param TreeNode $root
 * @return TreeNode
 */
function invertTree($root) {
	//左右数对调

	if($root == null) return null;

	list($root->left,$root->right) = [$this->invertTree($root->right),$this->invertTree($root->left)];

	return $root;
}