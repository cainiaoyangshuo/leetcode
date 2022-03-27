<?php
/**
 *
 * 定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

	示例:

	给定 1->2->3->4, 你应该返回 2->1->4->3.


 * @date 2020/7/8
 */

class Solution
{
    /**
	 * 方法一：递归
		这个题目要求我们从第一个节点开始两两交换链表中的节点，且要真正的交换节点。

		算法：

		从链表的头节点 head 开始递归。
		每次递归都负责交换一对节点。由 firstNode 和 secondNode 表示要交换的两个节点。
		下一次递归则是传递的是下一对需要交换的节点。若链表中还有节点，则继续递归。
		交换了两个节点以后，返回 secondNode，因为它是交换后的新头。
		在所有节点交换完成以后，我们返回交换后的头，实际上是原始链表的第二个节点。
	 * @param ListNode $head
	 * @return ListNode
	 */
    function swapPairs($head)
    {
        if (empty($head) || empty($head->next)) {
            return $head;
        }

        $second = $head->next;
        $head->next = $this->swapPairs($second->next);
        $second->next = $head;

        return $second;
    }

    /**
     * 二：迭代
     * @param $head
     */
    function swapPairs1($head)
    {
        if (empty($head) || empty($head->next)) {
            return $head;
        }

        // 虚拟头结点
        $dummyHead = new ListNode(-1);

        $dummyHead->next = $head;
        $cur = $dummyHead;

        while ($cur->next !== null && $cur->next->next !== null) {
            $a = $cur->next;
            $b = $cur->next->next;

            $cur->next = $b;
            $a->next = $b->next;
            $b->next = $a;

            $cur = $cur->next->next;
        }

        return $dummyHead->next;
    }
}
