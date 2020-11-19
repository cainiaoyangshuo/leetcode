<?php
/**
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	示例:

	给定的有序链表： [-10, -3, 0, 5, 9],

	一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

	 0
	/ \
  -3   9
  /   /
-10  5


 * @date 2020/8/31
 */

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution
{
    /**
	 * 标签：快慢指针
	 * 只需要找到他的中间结点，让他成为树的根节点，中间结点前面的就是根节点左子树的所有节点，
	 * 中间节点后面的就是根节点右子树的所有节点，然后使用递归的方式再分别对左右子树进行相同的操作……

	 * @param ListNode $head
	 * @return TreeNode
	 */
    function sortedListToBST($head)
    {
        // 边界条件
        if (empty($head)) {
            return null;
        }
        if (empty($head->next)) {
            return new TreeNode($head->val);
        }

        $slow = $head;
        $fast = $head;
        $prev = null;

        while ($fast != null && !empty($fast->next)) {
            $prev = $slow;
            $slow = $slow->next;
            $fast = $fast->next->next;
        }

        // 链表断开为两部分，前半是node左子树，后半是右子树
        $prev->next = null;
        // 当前节点
        $node = new TreeNode($slow->val);
        // head到prev，是左子树的节点
        $node->left = $this->sortedListToBST($head);
        // slow->next到链表末尾，是右子树的节点
        $node->right = $this->sortedListToBST($slow->next);

        return $node;
    }
}
