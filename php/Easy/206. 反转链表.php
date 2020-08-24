<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
require_once '/Users/yangshuo/git/leetcode/php/List/ListNode.php';
require_once '/Users/yangshuo/git/leetcode/php/List/List.php';
class Solution {

	/**
	 * 迭代
	 * 在遍历列表时，将当前节点的 next 指针改为指向前一个元素。
	 * 由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。
	 * 在更改引用之前，还需要另一个指针来存储下一个节点。不要忘记在最后返回新的头引用！
	 * @param ListNode $head
	 * @return ListNode
	 */
	function reverseList($head) {
		$prev = null;
		$cur = $head;

		while (!empty($cur)) {
			$temp = $cur->next;
			$cur->next = $prev;
			$prev = $cur;
			$cur = $temp;
		}

		return $prev;
	}

	/**
	 * 递归
	 *
	 * @param ListNode $head
	 * @return ListNode
	 */
	function reverseList1($head) {

	}
}

function test() {
	$arr = [1,2,3,4,5];

	// 创建链表
	$list = new LinkedList();
	foreach ($arr as $v) {
		$list->addNodeFromHead($v);
	}

	//遍历链表
//while (!empty($list->head)) {
//	print_r($list->head->val);
//	$list->head = $list->head->next;
//}

	$res = new Solution();
	$s = $res->reverseList($list->head);

	while (!empty($s)) {
		print_r($s->val);
		$s = $s->next;
	}
}

test();