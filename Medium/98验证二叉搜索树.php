<?php
/**
 * @date 2020/5/5
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
2
/ \
1   3
输出: true
示例 2:

输入:
5
/ \
1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value)
    {
        $this->val = $value;
    }
}

class Solution98
{
    protected $pre = PHP_INT_MIN;

    /**
     * 方法一：中序遍历
     * 中序遍历是二叉树的一种遍历方式，它先遍历左子树，再遍历根节点，最后遍历右子树。
     * 而我们二叉搜索树保证了左子树的节点的值均小于根节点的值，根节点的值均小于右子树的值，
     * 因此中序遍历以后得到的序列一定是升序序列。
     * @param TreeNode $root
     * @return Boolean
     */
    function isValidBST(TreeNode $root)
    {
        $a = PHP_INT_MIN;
        return $this->inOrder($root, $a);
    }

    function inOrder($root, &$a)
    {
        if ($root != null) {
            if (!$this->inOrder($root->left, $a)) {
                return false;
            }
            if ($a < $root->val) {
                $a = $root->val;
            } else {
                return false;
            }
            if (!$this->inOrder($root->right, $a)) {
                return false;
            }
        }
        return true;
    }
}
