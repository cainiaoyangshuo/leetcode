<?php
/**
 * @date 2020/6/19
 */

/**
 * 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 示例:
 * LRUCache cache = new LRUCache( 2 //缓存容量  );
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得关键字 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得关键字 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4

 * Class LRUCache
 *
 * 用哈希表和队列维护缓存中的键值对
 * 这样以来，我们首先使用哈希表进行定位，找出缓存项在双向链表中的位置，随后将其移动到双向链表的头部，
 * 即可在 O(1)O(1) 的时间内完成 get 或者 put 操作。具体的方法如下：
 * 对于 get 操作，首先判断 key 是否存在：
 * 如果 key 不存在，则返回 -1−1；
 * 如果 key 存在，则 key 对应的节点是最近被使用的节点。通过哈希表定位到该节点在双向链表中的位置，
 * 并将其移动到双向链表的头部，最后返回该节点的值。

 * 对于 put 操作，首先判断 key 是否存在：
 * 如果 key 不存在，使用 key 和 value 创建一个新的节点，在双向链表的头部添加该节点，并将 key 和该节点添加进哈希表中。
 * 然后判断双向链表的节点数是否超出容量，如果超出容量，则删除双向链表的尾部节点，并删除哈希表中对应的项；
 * 如果 key 存在，则与 get 操作类似，先通过哈希表定位，再将对应的节点的值更新为 value，并将该节点移到双向链表的头部。

 */
class LRUCache {
	/**
	 * @param Integer $capacity
	 */
	function __construct($capacity) {
		$this->used = [];   //k-v队列
		$this->hash = [];
		$this->capacity = $capacity;
	}

	public function keys()
	{
		return $this->used;
	}

	public function hash()
	{
		return $this->hash;
	}

	/**
	 * @param Integer $key
	 * @return Integer
	 */
	function get($key) {
		if (isset($this->hash[$key])) {
			$index = array_search($key, $this->used);
			unset($this->used[$index]);
			$this->used[] = $key;
			print_r($this->hash[$key]);
			return $this->hash[$key];
		}

		return -1;
	}

	/**
	 * @param Integer $key
	 * @param Integer $value
	 * @return NULL
	 */
	function put($key, $value) {
		if (isset($this->hash[$key])) {
			$index = array_search($key, $this->used);
			unset($this->used[$index]);
			$this->used[] = $key;
			$this->hash[$key] = $value;
		} else {
			if (count($this->used) == $this->capacity) {
				$headKey = array_shift($this->used);
				unset($this->hash[$headKey]);
			}
			$this->hash[$key] = $value;
			$this->used[] = $key;
		}
		return 'ok';
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * $obj = LRUCache($capacity);
 * $ret_1 = $obj->get($key);
 * $obj->put($key, $value);
 */

$obj = new LRUCache(2);

$obj->put(1,1);
$obj->put(2,2);
$obj->get(1);
$obj->put(3,3);

//var_dump($obj->keys(), $obj->hash());exit;
$obj->get(2);
$obj->put(4,4);
$obj->get(1);
$obj->get(3);
$obj->get(4);
