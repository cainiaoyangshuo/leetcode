<?php
/**
 * @date 2018/9/15
 * @author: yangshuo5@ucfgroup.com
 */

/**********************************************************************************
 *
 * You are given two linked lists representing two non-negative numbers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 *
 **********************************************************************************/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {

	/**
	 * @param ListNode $l1
	 * @param ListNode $l2
	 * @return ListNode
	 */
	function addTwoNumbers($l1, $l2) {
		if($l1==null){
			return $l2; //递归结束条件
		}

		if($l2==null){
			return $l1; //递归结束条件
		}

		$l1->next=$this->addTwoNumbers($l1->next,$l2->next);    //求出剩余链表的和，放入l1链表
		$l1->val=$l1->val+$l2->val;         //计算2个链表第一个元素和，放入l1链表
		$t=$l1->next;
		if($l1->val>=10){       //处理第一个元素相加大于等于10，产生进位的情况
			$l1->val=$l1->val%10;
			if($t==null){
				$l1->next=new ListNode(1);
			}else{
				while($t!=null){
					$t->val+=1;
					if($t->val<10){
						break;
					}else{
						$t->val=0;
						if($t->next==null){
							$t->next=new ListNode(1);
							break;
						}
					}
					$t=$t->next;
				}

			}
		}

		return $l1;     //最后将l1返回
	}


	function addTwoNumbersWhile($l1, $l2) {
		$sum=0;
		$step=0;    //代表进位
		$dummyHead=new ListNode(0);
		$cur=$dummyHead;
		while($l1!=null&&$l2!=null){
			$sum=$l1->val+$l2->val+$step;
			if($sum>=10){
				$cur->next=new ListNode($sum%10);
				$step=1;
			}else{
				$cur->next=new ListNode($sum);
				$step=0;
			}
			$cur=$cur->next;
			$l1=$l1->next;
			$l2=$l2->next;
		}

		while($l1!=null){
			$sum=$l1->val+$step;
			if($sum>=10){
				$cur->next=new ListNode($sum%10);
				$step=1;
			}else{
				$l1->val=$sum;
				$cur->next=$l1;
				$step=0;
				break;  //已经没有进位，提前结束循环
			}
			$cur=$cur->next;
			$l1=$l1->next;
		}

		while($l2!=null){
			$sum=$l2->val+$step;
			if($sum>=10){
				$cur->next=new ListNode($sum%10);
				$step=1;
			}else{
				$l2->val=$sum;
				$cur->next=$l2;
				$step=0;
				break;  //已经没有进位，提前结束循环
			}
			$cur=$cur->next;
			$l2=$l2->next;
		}

		if($step==1){
			$cur->next=new ListNode(1);     //表明还有进位
		}
		return $dummyHead->next;
	}
}


