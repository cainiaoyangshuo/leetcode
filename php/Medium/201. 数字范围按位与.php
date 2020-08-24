<?php
/**
 * 给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1: 

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * @date 2020/8/23
 */

class Solution {
	/**
	 * 我们通过右移，将两个数字压缩为它们的公共前缀。在迭代过程中，我们计算执行的右移操作数。
	 * 将得到的公共前缀左移相同的操作数得到结果。
	 * @param Integer $m
	 * @param Integer $n
	 * @return Integer
	 */
	function rangeBitwiseAnd($m, $n) {
		$shift = 0;
		// 找到公共前缀
		while ($m < $n) {
			$m >>= 1;
			$n >>= 1;
			++$shift;
		}

		return $m << $shift;
	}
	/**
	 * 复杂度分析
	时间复杂度：O(logn)。算法的时间复杂度取决于 m 和 n 的二进制位数，由于 m <= n，因此时间复杂度取决于 n 的二进制位数。
	空间复杂度：O(1)。我们只需要常数空间存放若干变量。
	 */
}