<?php
/**
 * 0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

	例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

	 

	示例 1：

	输入: n = 5, m = 3
	输出: 3
	示例 2：

	输入: n = 10, m = 17
	输出: 2
	 

	限制：

	1 <= n <= 10^5
	1 <= m <= 10^6

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * @date 2020/8/26
 */

class Solution {
	/**
	 * 数学方法 反推，res = (当前index + m) % 上一轮剩余数字的个数
	 * @param Integer $n
	 * @param Integer $m
	 * @return Integer
	 */
	function lastRemaining($n, $m) {
		$res = 0;
		// 最后一轮剩下两个人，所以从2开始反推
		for ($i=2; $i <= $n; $i++) {
			$res = ($res + $m) % $i;
		}

		return $res;
	}

}