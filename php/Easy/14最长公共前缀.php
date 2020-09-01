<?php

/**
 * 编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。

	示例 1:

	输入: ["flower","flow","flight"]
	输出: "fl"
	示例 2:

	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
	说明:

	所有输入只包含小写字母 a-z 。

 * @date 2020/6/10
 * @ide name PhpStorm
 */

class Solution {

	/**
	 * 显然，最长公共前缀的长度不会超过字符串数组中的最短字符串的长度。
	 * 用 minLength 表示字符串数组中的最短字符串的长度，
	 * 则可以在 [0,minLength] 的范围内通过二分查找得到最长公共前缀的长度。
	 * 每次取查找范围的中间值 mid，判断每个字符串的长度为 mid 的前缀是否相同，
	 * 如果相同则最长公共前缀的长度一定大于或等于 mid，如果不相同则最长公共前缀的长度一定小于 mid，
	 * 通过上述方式将查找范围缩小一半，直到得到最长公共前缀的长度。
	 * @param String[] $strs
	 * @return String
	 */
	function longestCommonPrefix($strs) {

	}

	/**
	 * 标签：滑动窗口
	 * 当字符串数组长度为 0 时则公共前缀为空，直接返回
	 * 令最长公共前缀 ans 的值为第一个字符串，进行初始化
	 * 遍历后面的字符串，依次将其与 ans 进行比较，两两找出公共前缀，最终结果即为最长公共前缀
	 * 如果查找过程中出现了 ans 为空的情况，则公共前缀不存在直接返回
	 * 时间复杂度：O(s)O(s)，s 为所有字符串的长度之和
	 */
	function two($strs) {
		$len = count($strs);
		if ($len == 0) {
			return '';
		}

		$ans = $strs[0];
		for ($i=1; $i<$len; $i++) {
			$j = 0;
			// j为公共前缀的长度
			for (; $j < strlen($ans) && $j < strlen($strs[$i]); $j++) {
				if ($ans[$j] != $strs[$i][$j]) {
					break;
				}
			}
			$ans = substr($ans, 0, $j);
			if ($ans == '') {
				return $ans;
			}
		}

		return $ans;
    }

}

$str = ["flower","flow","flight"];

