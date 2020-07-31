<?php
/**
 * @date 2020/7/15
 */

class Solution {

	public static $funcs = '';

	public function __construct() {



	}

	public function __call( $name, $arguments ) {
		// TODO: Implement __call() method.
		$ref = new ReflectionClass(__CLASS__);
		$funcs = $ref->getMethods();
		$func = $funcs[count($funcs)-1]->name;
		return $this->$func($arguments);
	}

	/**
	 *
	 * 1. 用变量max记录前一个子组合的最大值；用变量gMax记录全局最大值
	 * 2. 遍历数组，比较max，如果大于0，累加；否则抛弃前面的结果，将max置为当前值
	 * 3. 比较完后将最大值赋给gMax，遍历完gMax即为最大子序列之和
	 * @param Integer[] $nums
	 * @return Integer
	 */
	 public function maxSubArray($nums) {
		if (empty($nums)) {
			return false;
		}

		$max = 0;     // 前一个子组合的最大值
		$gMax = $nums[0];    // 全局最大值

		$len = count($nums);

		for ($i=0;$i<$len;$i++) {
			if ($max > 0) {
				// 如果前面子组合大于0，累加
				$max += $nums[$i];

			} else {
				// 小于0则抛弃前面的结果，从当前继续
				$max = $nums[$i];
			}
			// 比较全局最大值和前一个子组合的最大值，大的为新的全局最大值
			$gMax = max($gMax, $max);
		}

		return $gMax;
	}
}

$obj = new Solution();
$nums = [1];
$res = $obj->test($nums);
var_dump($res);