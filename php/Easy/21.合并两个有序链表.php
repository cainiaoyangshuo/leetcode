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
class Solution
{
    /**
     * 递归
     * 如果l1或l2为空链表，那么返回另一个非空链表
     * 否则，判断哪个链表头结点的值更小，然后递归地决定下一个添加到结果里的节点。如果两个链表有一个为空，递归结束。
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function mergeTwoLists($l1, $l2)
    {
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
	 * 思路
	 * 我们可以用迭代的方法来实现上述算法。
	 * 当 l1 和 l2 都不是空链表时，判断 l1 和 l2 哪一个链表的头节点的值更小，将较小值的节点添加到结果里，
	 * 当一个节点被添加到结果里之后，将对应链表中的节点向后移一位。


	 * @param ListNode $l1
	 * @param ListNode $l2
	 * @return ListNode
	 */
    function m2($l1, $l2)
    {
        $preHead = new ListNode();

        $prev = $preHead;

        while ($l2 != null && $l1 != null) {
            if ($l1->val <= $l2->val) {
                $prev->next = $l1;
                $l1 = $l1->next;
            } else {
                $prev->next = $l2;
                $l2 = $l2->next;
            }

            $prev = $prev->next;
        }
        // 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
        $prev->next = $l1 == null ? $l2 : $l1;

        // 为什么是return prehead-next呢？ 还没搞懂。
        return $preHead->next;
    }
}
