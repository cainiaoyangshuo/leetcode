<?php
/**
 * 现一个 MapSum 类里的两个方法，insert 和 sum。

	对于方法 insert，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。

	对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

	示例 1:

	输入: insert("apple", 3), 输出: Null
	输入: sum("ap"), 输出: 3
	输入: insert("app", 2), 输出: Null
	输入: sum("ap"), 输出: 5

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/map-sum-pairs
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * @date 2020/8/23
 */

class MapSum {
	/**
	 * Initialize your data structure here.
	 */
	function __construct() {

	}

	/**
	 * @param String $key
	 * @param Integer $val
	 * @return NULL
	 */
	function insert($key, $val) {

	}

	/**
	 * @param String $prefix
	 * @return Integer
	 */
	function sum($prefix) {

	}
}

/**
 * Your MapSum object will be instantiated and called as such:
 * $obj = MapSum();
 * $obj->insert($key, $val);
 * $ret_2 = $obj->sum($prefix);
 */