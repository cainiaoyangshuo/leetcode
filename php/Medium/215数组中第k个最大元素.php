<?php
/**
 * @date 2020/6/10
 */

class Solution {

	/**
	 * 维护一个size为k的最小堆
	 * @param Integer[] $nums
	 * @param Integer $k
	 * @return Integer
	 */
	function findKthLargest($nums, $k) {
		$heap = new SplMinHeap();
		$len = count($nums);

		for ($i=0;$i<$len;$i++) {
			// 堆里最多有k个元素
			if ($heap->count() < $k) {
				$heap->insert($nums[$i]);
			} elseif ($nums[$i] > $heap->top()) {
				// 注意 只有最小堆元素到k个时才进行比较
				$heap->extract();
				$heap->insert($nums[$i]);
			}
		}

		return $heap->top();
	}
}