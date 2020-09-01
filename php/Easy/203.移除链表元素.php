<?php
/**
 * 删除链表中等于给定值 val 的所有节点。

	示例:

	输入: 1->2->6->3->4->5->6, val = 6
	输出: 1->2->3->4->5
 * @date 2020/8/23
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
	 * 哨兵节点 虚拟节点。
	 * 头结点和中间节点的删除不同
	 *
	 * @param ListNode $head
	 * @param Integer $val
	 * @return ListNode
	 */
	function removeElements($head, $val) {
		$sentinal = new ListNode();
		// 哨兵节点next指向表头
		$sentinal->next = $head;

		// 初始化一个前继节点，一个当前节点
		$prev = $sentinal;
		$cur = $head;

		while ($cur != null) {
			if ($cur->val == $val) {
				$prev->next = $cur->next;
			} else {
				$prev = $cur;
			}
			$cur = $cur->next;
		}

		return $sentinal->next;
	}
}