<?php
/**
 * 找到两个单链表相交的起始节点。
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

class Solution {
	/**
	 * 我们通常做这种题的思路是设定两个指针分别指向两个链表头部，一起向前走直到其中一个到达末端，
	 * 另一个与末端距离则是两链表的 长度差。再通过长链表指针先走的方式消除长度差，最终两链表即可同时走到相交点。


	 * @param ListNode $headA
	 * @param ListNode $headB
	 * @return ListNode
	 */
	function getIntersectionNode($headA, $headB) {
		if (empty($headA) || empty($headB)) {
			return null;
		}

		$pA = $headA;
		$pB = $headB;

		while ($pA !== $pB) {
			// 短的链表到头了重新指向另一个表头

			$pA = $pA == null ? $headB : $pA->next;
			$pB = $pB == null ? $headA : $pB->next;
		}

		return $pA;
	}
}