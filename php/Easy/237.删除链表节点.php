<?php
/**
 * 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

	现有一个链表 -- head = [4,5,1,9]，它可以表示为:

	示例 1：

	输入：head = [4,5,1,9], node = 5
	输出：[4,1,9]
	解释：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
	示例 2：

	输入：head = [4,5,1,9], node = 1
	输出：[4,5,9]
	解释：给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

	提示：

	链表至少包含两个节点。
	链表中所有节点的值都是唯一的。
	给定的节点为非末尾节点并且一定是链表中的一个有效节点。
	不要从你的函数中返回任何结果。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * @date 2020/9/1
 */

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */

class Solution
{
    /**
     * 已知要删除单节点指针node，将node后一个结点的值赋给node，再将node与node下下一个结点链接。
     * @param ListNode $node
     * @return
     */
    function deleteNode($node)
    {
        //将node后一个结点的值赋给node
        $node->val = $node->next->val;
        //再将node与node下下一个结点链接
        $node->next = $node->next->next;
    }
}
