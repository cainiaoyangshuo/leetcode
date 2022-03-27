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
class Solution
{
    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    function levelOrder($root)
    {
        if (empty($root)) {
            return [];
        }

        $res = [];
        $queue = [];

        array_push($queue, $root);
        $level = 0;
        while ($count = count($queue)) {
            for ($i = 0; $i < $count; $i++) {
                $node = array_shift($queue);
                $res[$level][] = $node->val;

                if ($node->left) {
                    $queue[] = $node->left;
                }

                if ($node->right) {
                    $queue[] = $node->right;
                }
            }
            $level++;
        }

        return $res;
    }
}
