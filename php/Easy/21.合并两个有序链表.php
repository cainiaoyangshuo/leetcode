<?php
/**
 * @date 2020/7/15
 *
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
class Solution {

	/**
	 * 递归
	 * 如果l1或l2为空链表，那么返回另一个非空链表
	 * 否则，判断哪个链表头结点的值更小，然后递归地决定下一个添加到结果里的节点。如果两个链表有一个为空，递归结束。
	 * @param ListNode $l1
	 * @param ListNode $l2
	 * @return ListNode
	 */
	function mergeTwoLists($l1, $l2) {

		if (empty($l1)) {
			return $l2;
		} elseif (empty($l2)) {
			return $l1;
		} elseif ($l1->val < $l2->val) {
			$l1->next = $this->mergeTwoLists($l1->next, $l2);
			return $l1;
		} else {
			$l2->next = $this->mergeTwoLists($l1, $l2->next);
			return $l2;
		}
	}

	/**
	 * 迭代
	 * @param ListNode $l1
	 * @param ListNode $l2
	 * @return ListNode
	 */
	function m2($l1, $l2)
	{

	}
}