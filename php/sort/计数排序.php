<?php
/**
 * @date 2020/6/21
 */

class Solution {

	public static function sort($nums)
	{
		if (empty($nums)) {
			return [];
		}

		$length = count($nums);
		$max = $nums[0];
		for ($i = 0; $i < $length; $i++) {
			if ($nums[$i] > $max) {
				$max = $nums[$i];
			}
		}
		// 申请定长数组，长度为max+1，因为初始是0
		// 不能是php默认数组，因为不会重排下标。
		$bucket = new SplFixedArray($max+1);

		// 把数组元素以key放到数组，value为出现次数
		foreach ($nums as $num) {
			if (empty($bucket[$num])) {
				$bucket[$num] = 1;
			} else {
				$bucket[$num] += 1;
			}
		}

		// 按value次数输出
		foreach ($bucket as $index => $re) {
			echo str_repeat($index, $re);
		}

		return 0;
	}

}

$arr = [4,1,1,2,3,2,3,3,3];
Solution::sort($arr);