<?php
/**
 * @date 2020/6/19
 */

/**
 * Class LRUCache
 *
 */
class LRUCache {
	/**
	 * @param Integer $capacity
	 */
	function __construct($capacity) {
		$this->list = [];
		$this->hash = [];
		$this->capacity = $capacity;
	}

	/**
	 * @param Integer $key
	 * @return Integer
	 */
	function get($key) {

	}

	/**
	 * @param Integer $key
	 * @param Integer $value
	 * @return NULL
	 */
	function put($key, $value) {

	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * $obj = LRUCache($capacity);
 * $ret_1 = $obj->get($key);
 * $obj->put($key, $value);
 */