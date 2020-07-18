<?php
/**
 * @date 2020/7/9
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
	 * 借助辅助栈全部压栈，然后弹出k个
	 * @param ListNode $head
	 * @param Integer $k
	 * @return Integer
	 */
	function kthToLast($head, $k) {
		$stack = [];

		while (!empty($head)) {
			$stack[] = $head->val;
			$head = $head->next;
		}

		for ($i=0; $i<$k; $i++) {
			$res = array_pop($stack);
		}

		return $res;
	}

	/**
	 * 双指针
	 * @param $head
	 * @param $k
	 */
	function f2($head, $k)
	{
		$fast = $head;
		$slow = $head;

		for ($i = 0; $i < $k-1; $i++) {
			if (!empty($fast->next)) {
				$fast = $fast->next;
			}
		}

		while (!empty($fast->next)) {
			$fast = $fast->next;
			$slow = $slow->next;
		}

		return $slow->val;
	}
}